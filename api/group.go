package api

import "github.com/gin-gonic/gin"

func (api API) routeGroup(g *gin.RouterGroup) {
	g.GET("", api.listGroup)
	g.POST("", api.addGroup)

	{
		g := g.Group("/:group-id")
		g.GET("", api.getGroup)

		g.PUT("", api.updateGroup)
		g.DELETE("", api.deleteGroup)

		g.POST("/join", api.joinGroup)
		g.POST("/quit", api.joinGroup)
	}

}
