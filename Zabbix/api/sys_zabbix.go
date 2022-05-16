package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	Zabbixesult "github.com/flipped-aurora/gin-vue-admin/server/plugin/Zabbix/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Zabbix/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type zabbixAPI struct {
}

//@Tags Plugins
//@Summary 获取主机列表
//@Security ZabbixHostListGet
//@Produce  application/json
//@Success 200 {string} string "{"code":0,"data":{},"msg":"操作成功"}"
//@Router /zabbix/getHostList
func (s *zabbixAPI) GetHostAll(c *gin.Context) {
	var resultMail Zabbixesult.ResultModel
	_ = c.ShouldBindJSON(&resultMail)
	err, gethostALl := service.GetHost()
	if err != nil {
		global.GVA_LOG.Error("调用失败!", zap.Error(err))
		response.FailWithMessage("调用失败", c)
	} else {
		response.OkWithData(gethostALl, c)
	}
}

//@Tags Plugins
//@Summary 进行Web状态码监控
//@Security ZabbixWebCodeMonitor
//@Produce  application/json
//@Success 200 {string} string "{"code":0,"data":{},"msg":"操作成功"}"
//@Router /zabbix/getwebcode
func (s *zabbixAPI) GetWebMonitorStatus(c *gin.Context) {
	var webStatusUrl Zabbixesult.GetWebCodeStatusModel
	_ = c.ShouldBindJSON(&webStatusUrl)
	webStatusUrlCode, err := service.WebMonitor(webStatusUrl.Url)
	if err != nil {
		global.GVA_LOG.Error("调用失败!", zap.Error(err))
		response.FailWithMessage("调用失败", c)
	} else {
		response.OkWithData(&webStatusUrlCode, c)
	}
}

//@Tags Plugins
//@Summary 获取主机列表
//@Security GetTiggerList
//@Produce  application/json
//@Success 200 {string} string "{"code":0,"data":{},"msg":"操作成功"}"
//@Router /zabbix/getHostList
func (s *zabbixAPI) GetTiggerList(c *gin.Context) {
	var resultMail Zabbixesult.ResultModel
	_ = c.ShouldBindJSON(&resultMail)
	getTiggerList, err := service.GetAlertTiggerList()
	if err != nil {
		global.GVA_LOG.Error("调用失败!", zap.Error(err))
		response.FailWithMessage("调用失败", c)
	} else {
		response.OkWithData(getTiggerList, c)
	}
}
