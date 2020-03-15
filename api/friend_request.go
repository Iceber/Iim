package api

import "github.com/gin-gonic/gin"

func (api API) routeFriendRequest(g *gin.RouterGroup) {
	g.GET("", api.listFriendRequest)
	g.POST("", api.requestFriend)
	g.POST("/approved", api.approveFriendRequest)
}
