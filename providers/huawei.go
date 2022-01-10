package providers

import (
	"github.com/gourds/hw_cdn_refresh/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	cdn "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v1/model"
	"github.com/wonderivan/logger"
)

func (c *HuaWei)Session() (SessionType,error){
	auth := global.NewCredentialsBuilder().
		WithAk(c.AccessKeyID).
		WithSk(c.AccessKeySecret).
		WithDomainId(c.DomainId).
		Build()
	client := cdn.NewCdnClient(
		cdn.CdnClientBuilder().
			WithEndpoint(c.EndPoint).
			WithCredential(auth).
			Build())
	return SessionType{huawei: client},nil
}

func (c *HuaWei)Refresh(client SessionType) (err error){

	request := &model.CreateRefreshTasksRequest{}
	enterpriseProjectIdRequest:= c.ProjectId
	request.EnterpriseProjectId = &enterpriseProjectIdRequest
	var listUrlsRefreshTask = config.UrlsList
	refreshTaskbody := &model.RefreshTaskRequestBody{
		Urls: listUrlsRefreshTask,
	}
	request.Body = &model.RefreshTaskRequest{
		RefreshTask: refreshTaskbody,
	}
	response, err := client.huawei.CreateRefreshTasks(request)
	if err == nil {
		logger.Info("%+v", response)
		logger.Info("刷新成功：%v", response.HttpStatusCode)
	} else {
		logger.Error(err)
	}
	return err
}

func (c *HuaWei)Preheating(client SessionType) (err error){
	request := &model.CreatePreheatingTasksRequest{}
	enterpriseProjectIdRequest:= c.ProjectId
	request.EnterpriseProjectId = &enterpriseProjectIdRequest
	var listUrlsPreheatingTask = config.UrlsList
	preheatingTaskbody := &model.PreheatingTaskRequestBody{
		Urls: listUrlsPreheatingTask,
	}
	request.Body = &model.PreheatingTaskRequest{
		PreheatingTask: preheatingTaskbody,
	}
	response, err := client.huawei.CreatePreheatingTasks(request)
	if err == nil {
		logger.Info("%+v", response)
		logger.Info("刷新成功：%v", response.HttpStatusCode)

	} else {
		logger.Error(err)
	}
	return err
}