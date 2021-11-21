package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
	"xyz/dbettkk/btc/container"
)

var r *gin.Engine

func Init() {
	container.InitContainer()
	r = gin.Default()
	r.Use(Cors())
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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v\n", err)
				log.Printf("Panic info is: %s\n", debug.Stack())
			}
		}()

		c.Next()
	}
}
