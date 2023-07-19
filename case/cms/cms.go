package cms

import (
	"lbswebui/public"
	"time"

	"github.com/tebeka/selenium"
)

const (
	SERVER = public.SERVER
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
	err = WaitAlert(wd, public.ACCEPT, "")
	if err != nil {
		return err
	}
	return nil
}

func WaitAlert(wd selenium.WebDriver, op, expect string) error {
	var IsAlert = func(wd selenium.WebDriver) (bool, error) {
		s, err := wd.AlertText()
		if err != nil {
			// 还没有弹窗出现
			return false, nil // 继续到下一个间隔尝试获取alert()
		}
		if expect != "" && s != expect {
			return false, nil
		}
		return true, nil
	}
	err := wd.WaitWithTimeoutAndInterval(IsAlert, 10*time.Second, 500*time.Millisecond)
	if err != nil {
		return err
	}
	// 执行操作
	if op == public.ACCEPT {
		err = wd.AcceptAlert()
	} else {
		err = wd.DismissAlert()
	}
	return err
}
