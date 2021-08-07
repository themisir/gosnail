# 🐌 GoSnail

Experimental express.js like HTTP router written in golang. This project was
just a demostration to check whether or not we could create "developer friendly"
API for http routing.

## What does it look like?

Take a look at this sample code..

```go
import (
  "fmt"

  "gosnail"
)

func main() {
  // Router handles incoming requests from Application or another routers
  router := gosnail.NewRouter()

  // Application is wrapper for http.Handler that composes and redirects
  // requests to router
  app := gosnail.NewApplication(router)

  // Simple GET / request handler
  //
  // Arguments:
  //   ctx  - is request context contains request and response details
  //   next - calling this function fill forward request to the next handler
  router.Get("/", func (ctx *gosnail.Context, next func()) {
    fmt.Fprintf(ctx.Response(), "Index page")
  })

  // Middleware that handles all requests (excluding previously handled ones)
  //
  // Middlewares are simple functions that can intercept request and response
  // data and has ability to modifying them.
  //
  // In this example middleware, it simply checks if the request has an header
  // named 'Authorization', if not it simply responds with 401 status code. If
  // there's one, it forwards request to next handler.
  router.Use(gosnail.Any, func (ctx *gosnail.Context, next func()) {
    if ctx.Request().Headers.Get("Authorization") != "" {
      ctx.Response().StatusCode = 401
      ctx.Response().End()
    } else {
      next()
    }
  })

  // Another simple GET /users handler.
  
  // This handler will get called only if previously registered middleware
  // forwards request to it.
  router.Get("/users", func (ctx *gosnail.Context, next func()) {
    fmt.Fprintf(ctx.Response(), "This is supposed to be user list...")
  })

  // Creates http server that listens for incoming requests to port 8080 and
  // forwards them to the router.
  //
  // This method has same effect as calling http.ListenAndServe(":8080", app)
  app.Listen(":8080")
}
```

## Why?

Because current state of routing libraries in golang is not much friendly for
newcomer developers to golang coming from much simplier environments like
node.js. Bla blah... To be honest, I was just bored and tried to do something
for fun.

**It's not likely that I'm going to continue to work on this repository. So
please don't use it.**

Also if you read this far, I wanna let you know that I hate golang's module /
package management system, and that's what frustrated me a lot during making of
this package. That's it, thanks for coming to my TED talk.
