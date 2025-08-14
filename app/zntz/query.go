package zntz

import (
	"context"
	"jtyl_feishu_mt/model"

	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

func QueryRecord(recordId string) model.ZntzRecord {
	req := larkbitable.NewBatchGetAppTableRecordReqBuilder().
		AppToken(config.AppToken).
		TableId(config.TableId.TableIdZntz).
		Body(larkbitable.NewBatchGetAppTableRecordReqBodyBuilder().
			RecordIds([]string{recordId}).
			Build()).
		Build()

	resp, _ := client.Bitable.V1.AppTableRecord.BatchGet(context.Background(), req)
	var record model.ZntzRecord
	if resp != nil && resp.Data != nil && len(resp.Data.Records) > 0 {
		record.LtemType, _ = resp.Data.Records[0].Fields["物品类型"].(string)
		record.CommonItemOutName, _ = resp.Data.Records[0].Fields["常用品领用名称"].(string)
		record.OutQuantity, _ = resp.Data.Records[0].Fields["领用数量"].(float64)
		record.CommonItemInName, _ = resp.Data.Records[0].Fields["常用品入库名称"].(string)
		record.InQuantity, _ = resp.Data.Records[0].Fields["入库数量"].(float64)
	}
	return record
}

func QueryInventory(tableId string, recordId string) float64 {
	req := larkbitable.NewBatchGetAppTableRecordReqBuilder().
		AppToken(config.AppToken).
		TableId(tableId).
		Body(larkbitable.NewBatchGetAppTableRecordReqBodyBuilder().
			RecordIds([]string{recordId}).
			Build()).
		Build()

	resp, _ := client.Bitable.V1.AppTableRecord.BatchGet(context.Background(), req)
	var inventory float64
	if resp != nil && resp.Data != nil && len(resp.Data.Records) > 0 {
		inventory, _ = resp.Data.Records[0].Fields["剩余库存"].(float64)
	}
	return inventory
}
