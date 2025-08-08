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
