package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Method      string
	Path        string
	MiddleWares []gin.HandlerFunc
	Handler     gin.HandlerFunc
}

type routing struct {
	serverAddress string
	routers       []Router
}

type Routers interface {
	Serve()
}

func Initialize(serverAddress string, routers []Router) Routers {
	return &routing{serverAddress: serverAddress, routers: routers}
}

func (r *routing) Serve() {

	router := gin.Default()

	//assign path and method for the requests
	for _, route := range r.routers {
		method := route.Method
		switch method {
		case http.MethodPost:
			route.MiddleWares = append(route.MiddleWares, route.Handler)
			router.POST(route.Path, route.MiddleWares...)
		case http.MethodGet:
			route.MiddleWares = append(route.MiddleWares, route.Handler)
			router.GET(route.Path, route.MiddleWares...)
		case http.MethodPut:
			route.MiddleWares = append(route.MiddleWares, route.Handler)
			router.PUT(route.Path, route.MiddleWares...)
		case http.MethodDelete:
			route.MiddleWares = append(route.MiddleWares, route.Handler)
			router.DELETE(route.Path, route.MiddleWares...)
		}

	}
	//load all templates

	router.LoadHTMLGlob("../views/*.html")
	router.Static("/style", "../views/css")
	router.Static("/script", "../views/js")
	router.Static("/fonts", "../views/fontawesome-free-6.2.0-web")
	router.Static("/assets", "../views/assets")
	router.Static("/favicon.ico", "../views/assets/fv.ico")

	router.Run(r.serverAddress)

}
