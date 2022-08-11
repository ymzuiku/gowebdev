package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ymzuiku/webdev/helper"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	app := gin.New()

	if os.Getenv("dev") == "true" {
		helper.Proxy(app, "9100", "example/client/main.go")
	}

	log.Printf("listen: http://127.0.0.1:8300")
	if err := app.Run(":8300"); err != nil {
		fmt.Println("rightos app run err: ", err)
	}
}
