package model

type ZntzConfig struct {
	AppId     string
	AppSecret string
	AppToken  string
	TableId   TableSet
}

type TableSet struct {
	TableIdZntz string
	TableIdLx   string
	TableIdCyp  string
	TableIdZdh  string
	TableIdZs   string
}

type ZntzRecord struct {
	LtemType string `json:"物品类型"`

	CommonItemOutName string  `json:"常用品领用名称"`
	OutQuantity       float64 `json:"领用数量"`
	CommonItemInName  string  `json:"常用品入库名称"`
	InQuantity        float64 `json:"入库数量"`
}
type ReqestData struct {
	RecordId string `json:"record_id"`
}
type InData struct {
	PurchaseType         string  `json:"purchase_type"`
	CommonItemInName     string  `json:"common_item_in_name"`
	CommonItemInQuantity float64 `json:"common_item_in_quantity"`
}
type UpData struct {
	TableId   string
	RecordId  string
	Inventory float64
}
