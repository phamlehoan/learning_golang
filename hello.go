package main

import (
	"time"

	"github.com/kataras/iris/v12"
)

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Thời gian nhận request: %s", time.Now())
	ctx.Next()
}

func main() {
	app := iris.New()
	app.Use(myMiddleware)
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("Hello world")
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
