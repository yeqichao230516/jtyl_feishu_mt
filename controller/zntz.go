package controller

import (
	"jtyl_feishu_mt/app/zntz"
	"jtyl_feishu_mt/model"
	"jtyl_feishu_mt/service"

	"github.com/gin-gonic/gin"
)

func ZntzOut(c *gin.Context) {
	var req model.ReqestData
	if err := c.ShouldBindJSON(&req); err != nil {
		service.HandleError(c, 400, "Invalid JSON format", err)
		return
	}
	record := zntz.QueryRecord(req.RecordId)

	var updata model.UpData
	updata.TableId = getTableID(record.LtemType)
	updata.RecordId = getRecordID(record.CommonItemOutName)

	if zntz.QueryInventory(updata.TableId, updata.RecordId) < record.OutQuantity {
		zntz.DeleteZntzRecord(req.RecordId)
		service.HandleError(c, 400, "库存不足,已删除申请记录", nil)
		return
	}

	updata.Inventory = zntz.QueryInventory(updata.TableId, updata.RecordId) - record.OutQuantity
	zntz.UpdataRecord(updata.Inventory, service.GetString("zntz_table_cyp"), updata.RecordId)
	service.HandleSuccess(c, updata)

}
func ZntzIn(c *gin.Context) {
	var req model.ReqestData
	if err := c.ShouldBindJSON(&req); err != nil {
		service.HandleError(c, 400, "Invalid JSON format", err)
		return
	}
	record := zntz.QueryRecord(req.RecordId)
	var updata model.UpData
	updata.TableId = getTableID(record.LtemType)
	updata.RecordId = getRecordID(record.CommonItemInName)

	updata.Inventory = zntz.QueryInventory(updata.TableId, updata.RecordId) + record.InQuantity
	zntz.UpdataRecord(updata.Inventory, service.GetString("zntz_table_cyp"), updata.RecordId)
	service.HandleSuccess(c, updata)

}

func ZntzDataReturn(c *gin.Context) {
	var req model.ReqestData
	if err := c.ShouldBindJSON(&req); err != nil {
		service.HandleError(c, 400, "Invalid JSON format", err)
		return
	}
	record := zntz.QueryRecord(req.RecordId)
	var updata model.UpData
	updata.TableId = getTableID(record.LtemType)
	if record.CommonItemInName != "" {
		updata.RecordId = getRecordID(record.CommonItemInName)
		updata.Inventory = zntz.QueryInventory(updata.TableId, updata.RecordId) - record.InQuantity
		zntz.UpdataRecord(updata.Inventory, updata.TableId, updata.RecordId)
		service.HandleSuccess(c, updata)
		return
	} else if record.CommonItemOutName != "" {
		updata.RecordId = getRecordID(record.CommonItemOutName)
		updata.Inventory = zntz.QueryInventory(updata.TableId, updata.RecordId) + record.OutQuantity
		zntz.UpdataRecord(updata.Inventory, updata.TableId, updata.RecordId)
		service.HandleSuccess(c, updata)
		return
	}
}

func getRecordID(name string) string {
	switch name {
	case "A4纸":
		return service.GetString("cyp_record_A4纸")
	case "中性笔":
		return service.GetString("cyp_record_中性笔")
	case "中性笔芯":
		return service.GetString("cyp_record_中性笔芯")
	case "打印机碳粉":
		return service.GetString("cyp_record_打印机碳粉")
	case "记号笔":
		return service.GetString("cyp_record_记号笔")
	case "洗衣粉":
		return service.GetString("cyp_record_洗衣粉")
	case "洗手液":
		return service.GetString("cyp_record_洗手液")
	case "工作日志本":
		return service.GetString("cyp_record_工作日志本")
	default:
		return ""
	}
}
func getTableID(name string) string {
	switch name {
	case "常用品":
		return service.GetString("zntz_table_cyp")
	default:
		return ""
	}
}
