# webdev

Use gopherjs run web-dev-server

## cli

- web-develop-server
- build go to js (use gopherjs), and copy public to dist
- open box: tailwind.css
- full stack in your gin server, proxy to gopherjs-server

## Dev

Run server and, please open `http://127.0.0.1:8000` in your browser

```sh
webdev <your-pkg>/main.go
```

If you need change port:

```sh
webdev <your-pkg>/main.go --port 5000
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

## Appointment

- assets dir: <your-go-mod>/public
- root html: <your-go-mod>/index.html
