package apis

import (
	"encoding/json"
	"strconv"
	"txsms/tools"

	"github.com/labstack/echo"
	"github.com/shesuyo/yunpian"
)

// Template  使用模板发送
func Template(c echo.Context) error {
	mobile := c.FormValue("mobile")
	tplID := c.FormValue("tpl_id")
	data := c.FormValue("tpl_string")

	tplIDInt, err := strconv.Atoi(tplID)

	if mobile == "" {
		return tools.ReturnJSONResult(c, false, "mobile参数缺失", nil)
	}
	if tplID == "" || err != nil {
		return tools.ReturnJSONResult(c, false, "tpl_id参数错误", nil)
	}
	if data == "" {
		return tools.ReturnJSONResult(c, false, "tpl_string参数缺失", nil)
	}

	// 解析JSON start
	var p map[string]interface{}
	json.Unmarshal([]byte(data), &p)
	var tempString string
	for k, v := range p { // 构建模板字符串
		str, ok := v.(string)
		if ok {
			tempString += "#" + k + "#=" + str
		} else {
			num, ok := v.(float64)
			if ok {
				tempString += "#" + k + "#=" + tools.FloatToString(num)
			}
		}
	}
	// 解析JSON end

	// 发送到云片
	_, e := tools.APIInstance.SmsTplSend(yunpian.SMSTplSendInfo{Tpl_ID: tplIDInt, Mobile: mobile, Tpl_Value: tempString})
	if e != nil {
		return tools.ReturnJSONResult(c, false, "", e)
	}
	return tools.ReturnJSONResult(c, true, "", nil)
}
