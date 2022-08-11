package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ymzuiku/webdev/helper"
)

func main() {
	l := len(os.Args)
	if l < 2 {
		fmt.Printf("args need like:")
		fmt.Printf("bom <your-pkg>/main.go")
		fmt.Printf("bom <your-pkg>/main.go --port 4000")
		fmt.Printf("bom <your-pkg>/main.go --out dist")
		panic("args error.")
	}

	var dist string
	port := "9000"
	gopherJsPort := "8080"

	for i, v := range os.Args {
		if v == "--port" {
			port = os.Args[i+1]
		}
		if v == "--jsport" {
			gopherJsPort = os.Args[i+1]
		}
		if v == "--out" {
			dist = os.Args[i+1]
		}
	}

	if dist == "" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		app := gin.New()

		helper.Proxy(app, gopherJsPort, os.Args[1])

		log.Printf("webdev listening: http://127.0.0.1:" + port)
		if err := app.Run(":" + port); err != nil {
			fmt.Println("rightos app run err: ", err)
		}
	} else {
		helper.Build(os.Args[1], dist)
	}
}
