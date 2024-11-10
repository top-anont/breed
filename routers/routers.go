package routers

import "github.com/gin-gonic/gin"

func InitRouter(route *gin.RouterGroup) {
	BreedRoute(route)
}
