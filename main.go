package main

import (
	"jtyl_feishu_mt/app/zntz"
	"jtyl_feishu_mt/core"
)

func main() {
	initializeSystem()
	core.RunServe()
}
func initializeSystem() {
	core.NewViper()
	core.NewLogger()

	zntz.InitConfig()
	zntz.NewClient()

}
