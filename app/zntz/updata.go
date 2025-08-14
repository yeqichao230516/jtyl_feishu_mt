package zntz

import (
	"context"
	"jtyl_feishu_mt/service"

	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

func UpdataRecord(num float64, tableId string, recordId string) {
	req := larkbitable.NewUpdateAppTableRecordReqBuilder().
		AppToken(config.AppToken).
		TableId(tableId).
		RecordId(recordId).
		AppTableRecord(larkbitable.NewAppTableRecordBuilder().
			Fields(map[string]interface{}{`剩余库存`: num}).
			Build()).
		Build()

	resp, _ := client.Bitable.V1.AppTableRecord.Update(context.Background(), req)
	if resp != nil && resp.Code != 0 {
		service.Logger.Errorf("更新记录失败: %v", resp.Msg)
		return
	}
}
