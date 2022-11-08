package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func before(ctx iris.Context) {
	ctx.Application().Logger().Infof("Người dùng đã request: %s", ctx.Request().RequestURI)
	// gọi method Next để chuyển sang middleware tiếp theo
	ctx.Next()
}

func after(ctx iris.Context) {
	result := ctx.Values().Get("main")
	ctx.Application().Logger().Infof("%s", result)
	ctx.Application().Logger().Infof("Xử lý xong request")
}

// Hàm xử lý chính cũng được coi là một middleware
func homeHandler(ctx iris.Context) {
	ctx.View("home/index.html")
	// Truyền biến main sang middleware tiếp theo
	ctx.Values().Set("main", "received from main handler")
	ctx.Next()
}

func listProductsHandler(ctx iris.Context) {
	// // Get query string ?name=
	// // C1:
	// name := ctx.FormValue("name")
	// // C2:
	name := ctx.URLParam("name")
	ctx.Application().Logger().Infof("Name: %s", name)
	ctx.ViewData("Name", name)
	ctx.View("products/index.html")
}

func showProductHandler(ctx iris.Context) {
	ctx.View("products/show.html")
}

func newProductHandler(ctx iris.Context) {
	ctx.View("products/new.html")
}

func createProductHandler(ctx iris.Context) {
	name := ctx.PostValue("name") // Get data from form (name="name")
	ctx.Application().Logger().Infof("Name: %s", name)
	listProductsURL := fmt.Sprintf("/products?name=%s", name)
	ctx.Redirect(listProductsURL, iris.StatusMovedPermanently)
}

func editProductHandler(ctx iris.Context) {
	ctx.View("products/edit.html")
}

func main() {
	app := iris.New()
	app.HandleDir("/assets", "./assets")
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", before, homeHandler, after)
	app.PartyFunc("/products", func(products iris.Party) {
		products.Get("/", listProductsHandler)
		products.Get("/{productId}", showProductHandler)
		products.Get("/new", newProductHandler) // Không bị mâu thuẫn với product/{productId} phía trên
		products.Get("/edit", editProductHandler)
		products.Post("/", createProductHandler)
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
