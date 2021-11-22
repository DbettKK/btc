package db

import "C"
import (
	"xyz/dbettkk/btc/biz/container"
)

func GetTokens() []string {
	var tokenTables []TokenTable
	var tokens []string
	container.C.DB.Select("token").Find(&tokenTables)
	for _, token := range tokenTables {
		tokens = append(tokens, token.Token)
	}
	return tokens
}
