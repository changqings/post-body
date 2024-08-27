package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

var (
	serverAddr = ":30001"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	app.POST("/", func(c *gin.Context) {

		b, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{"status": "error"})
			return
		}

		fmt.Printf("post body:\n%s\n", string(b))

		c.JSON(200, gin.H{"status": "ok"})
	})

	err := app.Run(serverAddr)
	if err != nil {
		panic(err)
	}

}
