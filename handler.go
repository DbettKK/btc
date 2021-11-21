package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xyz/dbettkk/btc/dal/api"
	"xyz/dbettkk/btc/handler"
)

func IndexHandler(c *gin.Context) {
	if !handler.VerifyToken(c.GetHeader("Token")) {
		c.String(http.StatusForbidden, "invalid token ")
		return
	}
	btcReq := &api.BtcRequest{
		Start: 1,
		Limit: 2,
		Convert: "USD",
	}

	resp := handler.NewDefaultHandler(c).GetBtc(btcReq)
	c.JSON(http.StatusOK, resp)
	//c.String(http.StatusOK, resp)
}
