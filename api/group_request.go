package api

func (api API) routeGroupRequest(g *gin.RouterGroup) {
	g.GET("", api.listGroupRequest)
	g.POST("", api.requestGroup)
}
