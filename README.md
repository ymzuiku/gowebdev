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
