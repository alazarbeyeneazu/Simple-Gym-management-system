package main

import (
	handler "github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/handlers/rest/gym"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/admin"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/authz"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/checkin"
	gymgoers "github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/gym_goers"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/user"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/pyment"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/storage/persistant"
	routers "github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/routes"
)

func main() {

	dbs := persistant.Init()
	auth := authz.InitService(dbs)
	// auth.InitatePermission()
	service := user.InitService(dbs)
	pymentService := pyment.InitService(dbs)
	gymgoersService := gymgoers.InitService(dbs)
	checkinuser := checkin.InitService(dbs)
	admin := admin.InitService(dbs)
	// admin.InitializeSuperAdmin(context.Background(), "0948398647", "passme@123")
	user := handler.Init(service, pymentService, gymgoersService, admin, auth, checkinuser)
	routes := user.StartRoutes()
	router := routers.Initialize(":8282", routes)
	router.Serve()
}
