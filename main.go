package main

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/gin-gonic/gin"
)

var (
	serverAddr = ":8080"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	app.POST("/", func(c *gin.Context) {
		headers := c.Request.Header
		fmt.Println("post headers:")
		for k, v := range headers {
			fmt.Printf("%s: %s\n", k, v)
		}

		b, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{"status": "error"})
			return
		}

		fmt.Printf("post body:\n%s\n", string(b))

		c.JSON(200, gin.H{"status": "ok"})
	})

	slog.Info("start web", "server addr", serverAddr)
	err := app.Run(serverAddr)
	if err != nil {
		panic(err)
	}

}
