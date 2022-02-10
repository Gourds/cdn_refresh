package providers

import (
	"bytes"
	"encoding/json"
	"github.com/gourds/cdn_refresh/config"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"net/http"
)

func (c *BaiShan)Session() (SessionType, error){
	return SessionType{baishan: c.AccessKeySecret}, nil
}

type bsParams struct {
	Kind string  `json:"type"`
	Urls []string `json:"urls"`
}

func (c *BaiShan)Refresh(client SessionType) (err error){
	logger.Info("BaiShan Refresh Job Beginning...")
	url := "https://cdn.api.baishan.com/v2/cache/refresh?token=" + client.baishan
	params := &bsParams{
		Kind: "url",
		Urls: config.UrlsList,
	}

	//var jsonStr = []byte()
	jsonStr, err := json.Marshal(params)
	if err != nil {
		logger.Error("Error Marshal json")
	}
	req, err:= http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		logger.Error(err)
	}
	defer resp.Body.Close()
	logger.Info("response Status:", resp.Status)
	logger.Info("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	logger.Info("response Body:", string(body))
	return
}

func (c *BaiShan)Preheating(client SessionType)(err error){
	logger.Info("BaiShan Preheating Job Beginning...")
	url := "https://cdn.api.baishan.com/v2/cache/prefetch?token=" + client.baishan
	//var jsonStr = []byte(`{"urls"": config.UrlsList}`)
	jsonStr, err := json.Marshal(map[string][]string{"urls":config.UrlsList})
	req, err:= http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		logger.Error(err)
	}
	defer resp.Body.Close()
	logger.Info("response Status:", resp.Status)
	logger.Info("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	logger.Info("response Body:", string(body))
	return
}