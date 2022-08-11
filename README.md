# webdev

Use gopherjs run web-dev-server

## cli

- 启动开发服务
- 提供 gin 代理到 gopherjs
- 编译工程至生产
- 开箱 tailwind.css 支持

## Dev

```sh
webdev <your-pkg>/main.go
```

change port:

```sh
webdev <your-pkg>/main.go --port 500
```

## Build

```sh
webdev <your-pkg>/main.go --out dist
```

## Full stack

Use in gin server

```go
package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ymzuiku/webdev/helper"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	app := gin.New()

	if os.Getenv("dev") == "true" {
	  helper.Proxy(app, "8080", "example/client/main.go")
  }

	log.Printf("listen: http://127.0.0.1:8300")
	if err := app.Run(":8300"); err != nil {
		fmt.Println("rightos app run err: ", err)
	}
}

```

## 约定

- 资源文件夹 public; html 中引入资源: /public/xxxx.js
- html: index.html
