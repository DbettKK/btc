package main

import (
	"github.com/gin-gonic/gin"
	"xyz/dbettkk/btc/container"
)

var r *gin.Engine

func Init() {
	container.InitContainer()
	r = gin.Default()
	{
		r.GET("/", IndexHandler)
	}
}

func main() {
	Init()
	err := r.Run(":8000")
	if err != nil {
		return 
	}
}

