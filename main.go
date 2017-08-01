package main

import (
	"fmt"

	"github.com/shesuyo/yunpian"
)

var yunpianKey = "b92ee0b369f4749409bdb76bd2d13776"
var api *yunpian.YunpianAPI

func main() {
	api := yunpian.NewYunpianAPI(yunpianKey)
	r, e := api.SmsSend(yunpian.SMSSendInfo{Mobile: "18648552460", Text: "啊哈哈哈哈"})
	fmt.Println("---------- Start ----------")
	fmt.Println(r)
	fmt.Println(e)
	fmt.Println("----------- End -----------")
	fmt.Println("---------- Start2 ----------")
	r2, e2 := api.SmsTplSend(yunpian.SMSTplSendInfo{Tpl_ID: 1744884, Mobile: "18648552460", Tpl_Value: "#patient_name#=王志成"})
	fmt.Println(r2)
	fmt.Println(e2)
	fmt.Println("----------- End2 -----------")
}
