package routes

import (
	"user-service/controllers"
	routes "user-service/routes/user"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	controller controllers.IUserControllerRegistry
	group      *gin.RouterGroup
}

type IRouterRegistry interface {
	Serve()
}

func NewRouterRegistry(controller controllers.IUserControllerRegistry, group *gin.RouterGroup) IRouterRegistry {
	return &Registry{controller: controller, group: group}
}

func (r *Registry) Serve() {
	r.userRoute().Run()
}

func (r *Registry) userRoute() routes.IUserRoute {
	return routes.NewUserRoute(r.controller, r.group)
}
