package main

import (
  "fmt"
  "github.com/valyala/fasthttp"
)

func main() {
  fasthttp.ListenAndServe(":3000", helloWorld)
}

func helloWorld(ctx *fasthttp.RequestCtx) {
  fmt.Fprintf(ctx, "Hello world (go)!")
  ctx.SetContentType("text/plain; charset=utf8")
}
