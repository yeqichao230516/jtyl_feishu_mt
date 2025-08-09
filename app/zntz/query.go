package zntz

import (
	"context"
	"encoding/json"
	"fmt"
	"jtyl_feishu_mt/core"

	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

func QueryRecord(record_id string) {
	req := larkbitable.NewBatchGetAppTableRecordReqBuilder().
		AppToken(config.AppToken).
		TableId(config.TableId.TableIdZntz).
		Body(larkbitable.NewBatchGetAppTableRecordReqBodyBuilder().
			RecordIds([]string{record_id}).
			Build()).
		Build()

	resp, err := client.Bitable.V1.AppTableRecord.BatchGet(context.Background(), req)
	if err != nil {
		core.Logger.Error("查询记录失败")
	}
	if resp == nil || resp.Data == nil {
		core.Logger.Error("响应数据为空")
		return
	}
	r, _ := json.Marshal(resp.Data.Records)
	fmt.Println(string(r))
}
