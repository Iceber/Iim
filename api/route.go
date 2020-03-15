package api

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/uuid"
)

func (api API) RegisterRoutes(r *gin.Engine) {
	r.Use(addRequestID)

	api.routeUser(r.Group("/users"))
	api.routeGroup(r.Group("/groups"))
	api.routeFriends(r.Group("/friends"))
}

func addRequestID(c *gin.Context) {
	if c.GetHeader(HdrRequestID) == "" {
		c.Request.Header.Set(HdrRequestID, uuid.NewV4().String())
	}
}
