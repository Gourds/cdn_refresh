package providers

import (
	"github.com/gourds/hw_cdn_refresh/config"
	cdn "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v1"
)


type CDNManager interface {
	Session()(SessionType, error)
	Refresh(SessionType) error
	Preheating(SessionType) error
}

type SessionType struct {
	huawei *cdn.CdnClient
}

type HuaWei struct {
	config.Config
}

func GetConfig() (data CDNManager) {
	data = &HuaWei{config.InputConfig}
	return
}