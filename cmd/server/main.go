package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/static/*.html")
	router.GET("", func(c *gin.Context) { c.HTML(200, "index.html", nil) })
	router.Run(":8080")
}
