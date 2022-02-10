package providers

import (
	"github.com/gourds/cdn_refresh/config"
	cdn "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v1"
	"github.com/wonderivan/logger"
)


type CDNManager interface {
	Session()(SessionType, error)
	Refresh(SessionType) error
	Preheating(SessionType) error
}

type SessionType struct {
	huawei *cdn.CdnClient
	aliyun string
	baishan string
}

type HuaWei struct {
	config.Config
}

type AliYun struct {
	config.Config
}

type BaiShan struct {
	config.Config
}

func GetConfig() (data CDNManager) {
	logger.Info("Now Use Platform: %v", config.InputConfig.Platform)
	switch config.InputConfig.Platform {
	case "HuaWei":
		data = &HuaWei{config.InputConfig}
	case "AliYun":
		data = &AliYun{config.InputConfig}
	case "BaiShan":
		data = &BaiShan{config.InputConfig}
	default:
		data = &HuaWei{config.InputConfig}
	}
	return
}