package controller

import (
	"jtyl_feishu_mt/app/zntz"
	"jtyl_feishu_mt/model"
	"jtyl_feishu_mt/service"

	"github.com/gin-gonic/gin"
)

func ZntzOut(c *gin.Context) {
	var req model.OutData
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}
	var updata model.UpData
	switch req.PurchaseType {
	case "常用品":
		switch req.CommonItemOutName {
		case "A4纸":
			updata.RecordId = service.GetString("record_id_A4纸")
		case "中性笔":
			updata.RecordId = service.GetString("record_id_中性笔")
		case "中性笔芯":
			updata.RecordId = service.GetString("record_id_中性笔芯")
		case "打印机碳粉":
			updata.RecordId = service.GetString("record_id_打印机碳粉")
		case "记号笔":
			updata.RecordId = service.GetString("record_id_记号笔")
		case "洗衣粉":
			updata.RecordId = service.GetString("record_id_洗衣粉")
		case "洗手液":
			updata.RecordId = service.GetString("record_id_洗手液")
		case "工作日志本":
			updata.RecordId = service.GetString("record_id_工作日志本")
		}
		if zntz.QueryRecordInventoryFromCyp(updata.RecordId) < req.CommonItemOutQuantity {
			zntz.DeleteZntzRecord(req.RecordId)
			c.JSON(400, gin.H{"error": "库存不足"})
			return
		}
		updata.Inventory = zntz.QueryRecordInventoryFromCyp(updata.RecordId) - req.CommonItemOutQuantity
		zntz.UpdataRecord(updata.Inventory, service.GetString("table_id_cyp"), updata.RecordId)
		c.JSON(200, gin.H{"msg": "success", "data": updata})
		return
	}
	// case "自动化零配件":
	// 	updata.TableId = "tblxtyKeHorfu6SO"
	// case "注塑零配件":
	// 	updata.TableId = "tbl7rNQBQmnPwAvi"
	// }
}
func ZntzIn(c *gin.Context) {
	var req model.InData
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}
	var updata model.UpData
	switch req.PurchaseType {
	case "常用品":
		switch req.CommonItemInName {
		case "A4纸":
			updata.RecordId = service.GetString("record_id_A4纸")
		case "中性笔":
			updata.RecordId = service.GetString("record_id_中性笔")
		case "中性笔芯":
			updata.RecordId = service.GetString("record_id_中性笔芯")
		case "打印机碳粉":
			updata.RecordId = service.GetString("record_id_打印机碳粉")
		case "记号笔":
			updata.RecordId = service.GetString("record_id_记号笔")
		case "洗衣粉":
			updata.RecordId = service.GetString("record_id_洗衣粉")
		case "洗手液":
			updata.RecordId = service.GetString("record_id_洗手液")
		case "工作日志本":
			updata.RecordId = service.GetString("record_id_工作日志本")
		}
		updata.Inventory = zntz.QueryRecordInventoryFromCyp(updata.RecordId) + req.CommonItemInQuantity
		zntz.UpdataRecord(updata.Inventory, service.GetString("table_id_cyp"), updata.RecordId)
		c.JSON(200, gin.H{"msg": "success", "data": updata})
		return
	}
}
