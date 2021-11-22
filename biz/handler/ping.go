package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xyz/dbettkk/btc/biz/dal/api"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func IndexHandler(c *gin.Context) {
	if !VerifyToken(c.GetHeader("Token")) {
		c.String(http.StatusForbidden, "invalid token ")
		return
	}
	btcReq := &api.BtcRequest{
		Start:   1,
		Limit:   2,
		Convert: "USD",
	}

	resp := NewDefaultHandler(c).GetBtc(btcReq)
	c.JSON(http.StatusOK, resp)
	//c.String(http.StatusOK, resp)
}
