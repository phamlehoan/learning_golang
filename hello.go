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
	app.HandleDir("/assets", "./assets")
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.View("home/index.html")
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
