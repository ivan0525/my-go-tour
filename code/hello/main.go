package main

// 每个go程序都是由包组成的
// 程序运行的入口包是main
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello golang!",
	})
}

func main() {
	r := gin.Default()
	r.GET("/hello", sayHello)
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.Run(":9090")
}
