package main

import (
	"github.com/kataras/iris"
	"strconv"
)

func main() {
	a := iris.New()
	a.Get("/{number}", fizzbuzz)
	a.Run(iris.Addr(":8080"))
}

func fizzbuzz(ctx iris.Context) {
	if n, err := strconv.Atoi(ctx.Params().Get("number")); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Writef("Invalid number")
	} else if n%15 == 0 {
		ctx.Writef("fizzbuzz")
	} else if n%3 == 0 {
		ctx.Writef("fizz")
	} else if n%5 == 0 {
		ctx.Writef("buzz")
	} else {
		ctx.Writef("%d", n)
	}
}
