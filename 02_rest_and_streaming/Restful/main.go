package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ref:https://github.com/gin-gonic/gin
func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", getRequestWithParamInPath)
	// curl localhost:8080/welcomeMe?firstname=yiang\&lastname=peter
	router.GET("/welcomeMe", getRequestWithParamInParam)

	router.POST("/post", postRequestLogin)
	router.Run("localhost:8080")
}

func getRequestWithParamInPath(c *gin.Context) {
	// take a string and return a string ...
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s \n", name)
}

func getRequestWithParamInParam(c *gin.Context) {
	firstname := c.Request.URL.Query().Get("firstname")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	c.String(http.StatusOK, "Hello %s %s \n", firstname, lastname)
}

// Login is
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func postRequestLogin(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if json.User != "manu" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
