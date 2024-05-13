package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type key struct {
	Value string `json:"value"`
}

var keys = map[string]string{}

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
	router.PUT("/key-value-store/:key", postValue)
	router.GET("/key-value-store/:key", getValue)
	router.DELETE("/key-value-store/:key", deleteValue)

    router.Run(":8085")
}

func deleteValue(c *gin.Context) {
	key := c.Param("key")

	if _, exists := keys[key]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"doesExist": false, "error": "Key does not exist", "message": "Error in DELETE"})
	} else {
		delete(keys, key)
		c.JSON(http.StatusOK, gin.H{"doesExist": true, "message": "Deleted successfully"})
	}
}

func getValue(c *gin.Context) {
	key := c.Param("key")

	if _, exists := keys[key]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"doesExist": false, "error": "Key does not exist", "message": "Error in GET"})
	} else {
		value := keys[key]
		c.JSON(http.StatusOK, gin.H{"doesExist": true, "message": "Retrieved successfully", "value": value})
	}
}

func postValue(c *gin.Context) {
	var newKey key
	param := c.Param("key")

    if err := c.BindJSON(&newKey); err != nil {
        return
    } else if len(param) > 50 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key is too long", "message": "Error in PUT"})
	} else if newKey.Value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Value is missing", "message": "Error in PUT"})
	} else if _, exists := keys[param]; exists {
		keys[param] = newKey.Value
		c.JSON(http.StatusOK, gin.H{"message": "Updated successfully", "replaced": true})
	} else {
		keys[param] = newKey.Value
		c.JSON(http.StatusCreated, gin.H{"message": "Added successfully", "replaced": false})
	}
    
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

