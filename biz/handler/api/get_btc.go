package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"xyz/dbettkk/btc/biz/container"
	"xyz/dbettkk/btc/biz/dal/api"
	"xyz/dbettkk/btc/biz/dal/consts"
)

func SetParam(req *http.Request, btcReq *api.BtcRequest) {
	req.Header.Set("Accepts", "application/json")
	req.Header.Add(consts.ApiKeyHeader, consts.ApiKey)

	params := url.Values{}
	params.Add("start", strconv.Itoa(btcReq.Start))
	params.Add("limit", strconv.Itoa(btcReq.Limit))
	params.Add("convert", btcReq.Convert)

	req.URL.RawQuery = params.Encode()
}

func GetBtc(btcReq *api.BtcRequest) (*api.BtcResponse, error) {
	req, err := http.NewRequest("GET", consts.BtcURL, new(bytes.Buffer))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	SetParam(req, btcReq)

	resp, err := container.C.Api.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(resp.Status)

	respBody, _ := ioutil.ReadAll(resp.Body)
	respData := &api.BtcResponse{}

	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return respData, nil
}
