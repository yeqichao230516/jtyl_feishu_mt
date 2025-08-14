package zntz

import (
	"jtyl_feishu_mt/model"
	"jtyl_feishu_mt/service"

	lark "github.com/larksuite/oapi-sdk-go/v3"
)

var client *lark.Client
var config *model.ZntzConfig

func InitConfig() {
	config = &model.ZntzConfig{
		AppId:     service.GetString("jtyl_mt_app_id"),
		AppSecret: service.GetString("jtyl_mt_app_secret"),
		AppToken:  service.GetString("zntz_token"),
		TableId: model.TableSet{
			TableIdZntz: service.GetString("zntz_table_zntz"),
			TableIdCyp:  service.GetString("zntz_table_cyp"),
		},
	}
}

func NewClient() {
	client = lark.NewClient(config.AppId, config.AppSecret)
}
