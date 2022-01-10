package main

import (
	"github.com/gourds/hw_cdn_refresh/config"
	"github.com/gourds/hw_cdn_refresh/providers"
	"github.com/wonderivan/logger"
)




func main(){
	logger.Info(config.InputConfig)
	logger.Info(config.UrlsList)
	var data providers.CDNManager
	data = providers.GetConfig()
	client, err := data.Session()
	if err != nil{
		logger.Error(err)
	}
	switch config.InputConfig.Type {
	case "Refresh":
		data.Refresh(client)
	case "Preheating":
		data.Preheating(client)
	}
}
