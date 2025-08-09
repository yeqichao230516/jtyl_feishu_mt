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
		AppId:     service.GetString("feishu_app_id"),
		AppSecret: service.GetString("feishu_app_secret"),
		AppToken:  service.GetString("app_token"),
		TableId: model.TableSet{
			TableIdZntz: service.GetString("table_id_zntz"),
			TableIdLx:   service.GetString("table_id_lx"),
			TableIdCyp:  service.GetString("table_id_cyp"),
			TableIdZdh:  service.GetString("table_id_zdh"),
			TableIdZs:   service.GetString("table_id_zs"),
		},
	}
}

func NewClient() {
	client = lark.NewClient(config.AppId, config.AppSecret)
}
