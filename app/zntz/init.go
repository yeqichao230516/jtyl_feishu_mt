package zntz

import (
	"jtyl_feishu_mt/core"
	"jtyl_feishu_mt/model"

	lark "github.com/larksuite/oapi-sdk-go/v3"
)

var client *lark.Client
var config *model.ZntzConfig

func InitConfig() {
	config = &model.ZntzConfig{
		AppId:     core.GetString("feishu_app_id"),
		AppSecret: core.GetString("feishu_app_secret"),
		AppToken:  core.GetString("app_token"),
		TableId: model.TableSet{
			TableIdZntz: core.GetString("table_id_zntz"),
			TableIdLx:   core.GetString("table_id_lx"),
			TableIdCyp:  core.GetString("table_id_cyp"),
			TableIdZdh:  core.GetString("table_id_zdh"),
			TableIdZs:   core.GetString("table_id_zs"),
		},
	}
}

func NewClient() {
	client = lark.NewClient(config.AppId, config.AppSecret)
}
