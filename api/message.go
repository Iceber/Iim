package api

import "github.com/gin-gonic/gin"

func (api API) routeMessage(g *gin.RouterGroup) {
	g.GET("", api.getMessages)

	g.POST("", api.sendMessage)

	g.POST("/ack", api.ackMessage)
}
