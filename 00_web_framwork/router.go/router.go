package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// what is router, router is a device that routes the network traffic

const pingRouter = "/ping"

//NewRouter tells at certain path what to do
func NewRouter() (*gin.Engine, error) {
	r := gin.Default()
	r.GET(pingRouter, func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	return r, nil
}
