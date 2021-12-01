package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"xyz/dbettkk/btc/biz/container"
)

func main() {
	container.InitContainer()
	r := gin.Default()

	// waitFunc, mw := shutdown.New(10e9, 0)
	// r.Use(mw)
	// defer waitFunc()

	register(r)

	if err := r.Run(":8000"); err != nil {
		os.Exit(1)
	}
}
