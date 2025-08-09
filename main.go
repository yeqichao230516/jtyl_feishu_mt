package main

import (
	"jtyl_feishu_mt/app/zntz"
	"jtyl_feishu_mt/core"
	"jtyl_feishu_mt/service"
)

func main() {
	initializeSystem()
	core.RunServe()
}

func initializeSystem() {
	service.NewViper()
	service.NewLogger()

	zntz.InitConfig()
	zntz.NewClient()
}
