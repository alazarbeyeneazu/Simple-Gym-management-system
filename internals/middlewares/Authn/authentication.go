package authn

import (
	"context"
	"net/http"
	"strings"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	token "github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/middlewares/token"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/storage/persistant"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// var authorizationType = "bearer"

func Authorized(action string, userid uuid.UUID) bool {
	dbs := persistant.Init()
	user, err := dbs.GetUserById(context.Background(), userid)
	if err != nil {
		return false
	}
	role, err := dbs.GetAssignRoleByUserId(context.Background(), models.UserRole{
		UserId: user.ID,
	})
	if err != nil {
		return false
	}
	rolePermissions, err := dbs.GetRolesByName(context.Background(), models.Role{
		RoleName: role.RoleName,
	})
	if err != nil {
		return false
	}
	for _, permissions := range rolePermissions {
		perm, err := dbs.GetPermissionById(context.Background(), models.Permission{ID: permissions.PermissionID})
		if err != nil {
			continue
		} else {
			if perm.Action == action {
				return true
			}
		}
	}

	return false
}

func AuthenticatRequest() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		path := ctx.Request.URL
		method := ctx.Request.Method
		paths := strings.Split(path.Path, "/")

		maker, err := token.NewPastoMaker("01234567890123456789012345678912")
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:8282/view/login")
			ctx.AbortWithStatus(http.StatusTemporaryRedirect)
			return
		}
		authHeader, err := ctx.Cookie("Athorization")
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:8282/view/login")
			ctx.AbortWithStatus(http.StatusTemporaryRedirect)
			return
		}
		if authHeader == "" {
			ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:8282/view/login")
			ctx.AbortWithStatus(http.StatusTemporaryRedirect)
			return
		}

		pload, err := maker.VerifyToken(authHeader)

		if err != nil {
			ctx.SetCookie("error", "login session expired", 2, "/", "localhost", false, false)
			ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:8282/view/login")
			ctx.AbortWithStatus(http.StatusTemporaryRedirect)
			return
		}
		_, err = uuid.Parse(paths[len(paths)-1])

		if err != nil {

			if paths[len(paths)-1] == "users" && method == http.MethodGet {
				if !Authorized("list Users", pload.UserId) {
					ctx.Redirect(http.StatusTemporaryRedirect, "/view/dashboard")
					ctx.Abort()
					return
				}

			} else if paths[len(paths)-1] == "roles" && method == http.MethodGet {
				if !Authorized("View Roles", pload.UserId) {
					ctx.Redirect(http.StatusTemporaryRedirect, "/view/dashboard")
					ctx.Abort()
					return
				}
			} else if paths[len(paths)-1] == "gym-goers" && method == http.MethodGet {
				if !Authorized("List GymGoer", pload.UserId) {
					ctx.Redirect(http.StatusTemporaryRedirect, "/view/dashboard")
					ctx.Abort()
					return
				}
			} else if paths[len(paths)-1] == "payment" && method == http.MethodGet {
				if !Authorized("List Payment", pload.UserId) {
					ctx.Redirect(http.StatusTemporaryRedirect, "/view/dashboard")
					ctx.Abort()
					return
				}
			} else if paths[len(paths)-1] == "reportbydays" || paths[len(paths)-1] == "reports" && method == http.MethodGet {
				if !Authorized("view report", pload.UserId) {
					ctx.Redirect(http.StatusTemporaryRedirect, "/view/dashboard")
					ctx.Abort()
					return
				}
			}

		} else {

			if paths[len(paths)-2] == "users" && method == http.MethodGet {
				if !Authorized("Edit User", pload.UserId) {
					ctx.Redirect(http.StatusTemporaryRedirect, "/view/dashboard")
					return
				}

			} else if paths[len(paths)-2] == "gym-goers" && method == http.MethodGet {
				if !Authorized("CheckIn gymgoer", pload.UserId) {
					ctx.Redirect(http.StatusTemporaryRedirect, "/view/dashboard")
					return
				}

			} else if paths[len(paths)-2] == "delete" {
				if paths[len(paths)-3] == "gymgoers" && method == http.MethodGet {
					if !Authorized("Delete Gymgoer", pload.UserId) {
						ctx.Redirect(http.StatusTemporaryRedirect, "/view/dashboard")
						ctx.Abort()
						return
					}

				}
			} else if paths[len(paths)-2] == "payment" && method == http.MethodGet {
				if !Authorized("Edit Payment", pload.UserId) {
					ctx.Redirect(http.StatusTemporaryRedirect, "/view/dashboard")
					return
				}

			} else if paths[len(paths)-2] == "gym-goers-detail" && method == http.MethodGet {
				if !Authorized("View Detail of Gym-Goer", pload.UserId) {
					ctx.Redirect(http.StatusTemporaryRedirect, "/view/dashboard")
					ctx.Abort()
					return
				}

			}

		}
		ctx.Set("userid", pload.UserId)
		ctx.Next()

	}
}
