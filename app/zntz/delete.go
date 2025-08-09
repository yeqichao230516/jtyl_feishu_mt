package zntz

import (
	"context"
	"jtyl_feishu_mt/service"

	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

func DeleteZntzRecord(recordId string) {
	req := larkbitable.NewDeleteAppTableRecordReqBuilder().
		AppToken(config.AppToken).
		TableId(config.TableId.TableIdZntz).
		RecordId(recordId).
		Build()

	resp, _ := client.Bitable.V1.AppTableRecord.Delete(context.Background(), req)

	if resp != nil && resp.Code != 0 {
		service.Logger.Errorf("删除记录失败: %v", resp.Msg)
		return
	}
}
