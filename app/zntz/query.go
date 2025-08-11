package zntz

import (
	"context"
	"jtyl_feishu_mt/model"

	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

func QueryRecord(record_id string) model.ZntzRecord {
	req := larkbitable.NewBatchGetAppTableRecordReqBuilder().
		AppToken(config.AppToken).
		TableId(config.TableId.TableIdZntz).
		Body(larkbitable.NewBatchGetAppTableRecordReqBodyBuilder().
			RecordIds([]string{record_id}).
			Build()).
		Build()

	resp, _ := client.Bitable.V1.AppTableRecord.BatchGet(context.Background(), req)
	var record model.ZntzRecord
	if resp != nil && resp.Data != nil && len(resp.Data.Records) > 0 {
		record.PurchaseType, _ = resp.Data.Records[0].Fields["采购申请类型"].(string)

		record.CommonItemOutName, _ = resp.Data.Records[0].Fields["常用品领用名称"].(string)
		record.CommonItemOutQuantity, _ = resp.Data.Records[0].Fields["常用品领用数量"].(float64)
		record.CommonItemInName, _ = resp.Data.Records[0].Fields["常用品入库名称"].(string)
		record.CommonItemInQuantity, _ = resp.Data.Records[0].Fields["常用品入库数量"].(float64)

		// record.Fields.AutoPartsOutName, _ = resp.Data.Records[0].Fields["自动化零配件领用名称"].(string)
		// record.Fields.AutoPartsOutQuantity, _ = resp.Data.Records[0].Fields["自动化零配件领用数量"].(float64)
		// record.Fields.AutoPartsInName, _ = resp.Data.Records[0].Fields["自动化零配件入库名称"].(string)
		// record.Fields.AutoPartsInQuantity, _ = resp.Data.Records[0].Fields["自动化零配件入库数量"].(float64)

		// record.Fields.MoldingPartsOutName, _ = resp.Data.Records[0].Fields["注塑零配件领用名称"].(string)
		// record.Fields.MoldingPartsOutQuantity, _ = resp.Data.Records[0].Fields["注塑零配件领用数量"].(float64)
		// record.Fields.MoldingPartsInName, _ = resp.Data.Records[0].Fields["注塑零配件入库名称"].(string)
		// record.Fields.MoldingPartsInQuantity, _ = resp.Data.Records[0].Fields["注塑零配件入库数量"].(float64)

	}

	return record
}

func DataReturn(record_id string) {
	req := larkbitable.NewBatchGetAppTableRecordReqBuilder().
		AppToken(config.AppToken).
		TableId(config.TableId.TableIdZntz).
		Body(larkbitable.NewBatchGetAppTableRecordReqBodyBuilder().
			RecordIds([]string{record_id}).
			Build()).
		Build()

	resp, _ := client.Bitable.V1.AppTableRecord.BatchGet(context.Background(), req)
	if resp != nil && resp.Data != nil && len(resp.Data.Records) > 0 {
		resp.Data.Records[0].Fields["采购申请类型"] = "常用品"
	}
}
func QueryRecordInventoryFromCyp(record_id string) float64 {
	req := larkbitable.NewBatchGetAppTableRecordReqBuilder().
		AppToken(config.AppToken).
		TableId(config.TableId.TableIdCyp).
		Body(larkbitable.NewBatchGetAppTableRecordReqBodyBuilder().
			RecordIds([]string{record_id}).
			Build()).
		Build()

	resp, _ := client.Bitable.V1.AppTableRecord.BatchGet(context.Background(), req)
	var inventory float64
	if resp != nil && resp.Data != nil && len(resp.Data.Records) > 0 {
		inventory, _ = resp.Data.Records[0].Fields["剩余库存"].(float64)
	}
	return inventory
}
