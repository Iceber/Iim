package api

import "github.com/gin-gonic/gin"

func (api API) routeFriend(g *gin.RouterGroup) {
	g.GET("", api.listFriends)
	g.PUT("", api.updateFriend)

	api.routeFriendRequest(g.Group("requests"))

	{
		g := g.Group("/:friend-id")
		g.GET("", api.getFriend)
		g.DELETE("", api.deleteFriend)
	}
}
