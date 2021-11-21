package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	ctx *gin.Context
}

func NewDefaultHandler(c *gin.Context) *Handler {
	return &Handler{ctx: c}
}
