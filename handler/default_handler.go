package handler

import (
	"strings"
	api2 "xyz/dbettkk/btc/dal/api"
	"xyz/dbettkk/btc/dal/db"
	"xyz/dbettkk/btc/handler/api"
)

func VerifyToken(token string) bool {
	tokens := db.GetTokens()
	for _, t := range tokens {
		if strings.Compare(t, token) == 0 {
			return true
		}
	}
	return false
}

func (h *Handler) GetBtc(btcReq *api2.BtcRequest) *api2.BtcResponse {
	btcResp, err := api.GetBtc(btcReq)
	if err != nil {
		return nil
	}
	//resp, _ := json.Marshal(btcResp)
	return btcResp
}
