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
	PurchaseType string `json:"采购申请类型"`

	CommonItemOutName     string  `json:"常用品领用名称"`
	CommonItemOutQuantity float64 `json:"常用品领用数量"`
	CommonItemInName      string  `json:"常用品入库名称"`
	CommonItemInQuantity  float64 `json:"常用品入库数量"`
}
type OutData struct {
	RecordId              string  `json:"record_id"`
	PurchaseType          string  `json:"purchase_type"`
	CommonItemOutName     string  `json:"common_item_out_name"`
	CommonItemOutQuantity float64 `json:"common_item_out_quantity"`
}
type InData struct {
	PurchaseType         string  `json:"purchase_type"`
	CommonItemInName     string  `json:"common_item_in_name"`
	CommonItemInQuantity float64 `json:"common_item_in_quantity"`
}
type UpData struct {
	RecordId  string
	Inventory float64
}
