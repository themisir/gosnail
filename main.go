package main

import (
	"fmt"
)

func main() {
	router := NewRouter()
	app := NewApplication(router)

	router.Use(Any, func(ctx *Context, next func()) {
		fmt.Printf("---> %s %s\n", ctx.Request().Method, ctx.Request().URL)
		next()
		fmt.Printf("<--- %d %s\n", ctx.Response().StatusCode, ctx.Request().URL)
	})

	router.Get("/", func(ctx *Context, next func()) {
		fmt.Fprintf(ctx.Response(), "I got you!")
	})

	router2 := NewRouter()

	router2.Get("/auth/login", func(ctx *Context, next func()) {
		ctx.Response().Headers().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(ctx.Response(), `<form method="post" action="/auth/login?username=bob&password=pas$">
			<button type="submit">click me</button>
		</form>`)
	})

	router2.Post("/auth/login", func(ctx *Context, next func()) {
		fmt.Fprintf(
			ctx.Response(),
			"Username: %s, Password: %s",
			ctx.Request().Query("username"),
			ctx.Request().Query("password"),
		)
	})

	router.Use("/auth", router2.Next)

	router3 := NewRouter()

	router3.Use(Any, func(ctx *Context, next func()) {
		if ctx.Request().Query("auth") != "1" {
			ctx.Response().StatusCode = 401
			ctx.Response().End()
			return
		}

		next()
	})

	router3.Get("/token", func(ctx *Context, next func()) {
		fmt.Fprintf(ctx.Response(), "Token: ce62ff2ec91f4ce2935c8b3c36a42722")
	})

	router.Use(Any, router3.Next)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
