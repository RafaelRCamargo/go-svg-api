package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	. "github.com/lunux2008/xulu"
)

func Server(PORT string) {
	println("Starting server...")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<div>aaa</div>",
		})
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./static/favicon.ico")
	})

	port := ":8080"

	println("\nPORT:\n", PORT)

	if PORT != "" {
		port := PORT
		Use(port)
	}

	r.Run(port)

	println("Listing for requests at http://localhost/" + os.Getenv("PORT"))
}
