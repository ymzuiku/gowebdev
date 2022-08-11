package helper

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ymzuiku/webdev/pkg/execx"
	"github.com/ymzuiku/webdev/pkg/fsx"
)

func Proxy(app *gin.Engine, port, clienPath string) {
	InitBaseFiles()
	go func() {
		execx.Run(context.Background(), nil, "gopherjs", "serve", clienPath)
	}()
	go func() {
		if fsx.Exists("tailwind.config.js") {
			execx.Run(context.Background(), nil, "tailwindcss", "./tailwind.css", "-o", "./public/tailwind.css", "--watch")
		}
	}()

	app.Use(func(ctx *gin.Context) {
		remotePath := ctx.Request.URL.String()
		if strings.Contains(remotePath, "/v1") {
			return
		}
		if strings.Contains(remotePath, "/public") {
			ctx.File("." + remotePath)
			return
		}
		// html 文件
		if remotePath == "/" {
			html := fsx.LoadFile("index.html")
			ctx.Header("Content-Type", "text/html; charset=utf-8")
			ctx.Header("Cache-control", "no-cache")
			ctx.String(200, html)
			return
		}

		// 如果请求 gojs 文件, 做一些编译事件
		// if remotePath == "/main.go.js" {
		// 	//
		// }

		// 代理
		remote, err := url.Parse("http://127.0.0.1:" + port + remotePath)
		if err != nil {
			fmt.Println("proxy:err", err)
			ctx.Abort()
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = ctx.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = remote.Path
		}
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	})
}
