package main

import (
	"encoding/json"
	"fmt"

	"github.com/tebeka/selenium"
)

func AddCookie(wd selenium.WebDriver) {
	// 创建一个Cookie
	cookie := &selenium.Cookie{
		Name:   "session",
		Value:  "ba76cmfsfi9o5m987ts6eu24hc",
		Path:   "/",
		Domain: "192.168.190.128",
	}

	// 设置Cookie
	err := wd.AddCookie(cookie)
	if err != nil {
		fmt.Println("Failed to add cookie: ", err)
		return
	}
}

type certData struct {
	Cert     string `json:"cert"`
	Key      string `json:"key"`
	Password string `json:"password"`
}
type lbsmap map[string]certData

func main() {
	certMap := make(map[string]certData)
	certMap["c1"] = certData{
		Cert:     "D:\\File\\work\\task\\GAD测试\\验证扩展用法\\server2\\cert_rsa_2048_lbs-cs-2.com_4688d7aa-bd26-489c-91a1-c8ad21af8182.pem",
		Key:      "D:\\File\\work\\task\\GAD测试\\验证扩展用法\\server2\\cert_rsa_2048_lbs-cs-2.com_4688d7aa-bd26-489c-91a1-c8ad21af8182.key",
		Password: "",
	}
	rsp, err := json.Marshal(certMap)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(rsp))
	}

	lbstest := lbsmap{}
	err = json.Unmarshal(rsp, lbstest)
	if err != nil {
		fmt.Println(err)
	}
	// 运行测试用例
	// cms.Test_c1(wd)
}
