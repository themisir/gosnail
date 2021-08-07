package gosnail

import (
	"fmt"

	"github.com/themisir/gosnail"
)

func main() {
	router := gosnail.NewRouter()
	app := gosnail.NewApplication(router)

	router.Get("/", func(ctx *gosnail.Context, next func()) {
		fmt.Fprintf(ctx.Response(), "I got you! Param: %s", ctx.Request().Query("param"))
	})

	app.Listen(":8080")
}
