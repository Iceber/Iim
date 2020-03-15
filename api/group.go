package api

import "github.com/gin-gonic/gin"

func (api API) routeGroup(g *gin.RouterGroup) {
	g.GET("", api.listGroups)
	g.POST("", api.addGroup)

	{
		g := g.Group("/:group-id")
		g.GET("", api.getGroup)

		g.PUT("", api.updateGroup)
		g.DELETE("", api.deleteGroup)
	}
}
