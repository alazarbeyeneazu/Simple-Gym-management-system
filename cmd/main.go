package main

import (
	handler "github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/handlers/rest/gym"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/pyment"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/user"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/storage/persistant"
	routers "github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/routes"
)

func main() {
	dbs := persistant.Init()
	service := user.InitService(dbs)
	pymentService := pyment.InitService(dbs)
	user := handler.Init(service, pymentService)
	routes := user.StartRoutes()
	router := routers.Initialize(":8282", routes)
	router.Serve()
}
