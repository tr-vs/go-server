package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "I'm alive!!"})
}

func main() {
    router := gin.Default()
	router.GET("/ping", getPing)
	router.POST("/ping", postPing)
	router.POST("/ping/:name", postPingByName)
	router.GET("/ping/:name", getPingByName)
	router.POST("/echo", postEcho)
	router.GET("/echo", getEcho)

    router.Run(":8085")
}

func getEcho(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get Message Received"})
}

func postEcho(c *gin.Context) {
	msg := c.Query("msg")
	if msg == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": msg})
	}
}

func getPingByName(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method Not Allowed"})
}

func postPingByName(c *gin.Context) {
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{"message": "I'm alive, " + name + "!!"})
}

func postPing(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method Not Allowed"})
}

