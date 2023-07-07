package cms

import (
	"lbswebui/public"

	"github.com/tebeka/selenium"
)

const (
	SERVER = "https://192.168.190.128:60443"
)

// 删除所有证书
func DeleteAllCert(wd selenium.WebDriver, url string) error {
	var err error
	// 进入站点证书界面
	err = wd.Get(url)
	if err != nil {
		return err
	}
	// 勾选所有
	eles, err := wd.FindElements(selenium.ByName, "id[]")
	if err != nil {
		return err
	}
	for _, ele := range eles {
		ele.Click()
		if err != nil {
			return err
		}
	}
	// 删除所有
	del, err := wd.FindElement(selenium.ByID, "btn_del")
	if err != nil {
		return err
	}
	err = del.Click()
	if err != nil {
		return err
	}
	// 确认删除
	err = wd.AcceptAlert()
	if err != nil {
		return err
	}
	// 证书被其他服务引用，弹窗显示
	err = wd.AcceptAlert()
	if err != nil {
		// 有服务引用
	} else {
		// 没有服务引用
	}
	return nil
}

func SwitchToPage(url string) (selenium.WebDriver, error) {
	var err error
	wd, err := public.GetLogin()
	if err != nil {
		return nil, err
	}

	// 切换到证书请求
	err = wd.Get(url)
	if err != nil {
		return nil, err
	}
	return wd, nil
}
