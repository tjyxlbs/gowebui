package public

import (
	"time"

	"github.com/tebeka/selenium"
)

const (
	ACCEPT  = "accept"
	DISMISS = "DISMISS"
)

func TextEle(wd selenium.WebDriver, id, value string) error {
	ele, err := wd.FindElement(selenium.ByID, id)
	if err != nil {
		return err
	}
	err = ele.Clear()
	if err != nil {
		return err
	}
	err = ele.SendKeys(value)
	if err != nil {
		return err
	}
	return nil
}

func ClickByXpath(wd selenium.WebDriver, xpath string) error {
	// 下一步
	next, err := wd.FindElement(selenium.ByXPATH, xpath)
	if err != nil {
		return err
	}
	err = next.Click()
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
	if op == ACCEPT {
		err = wd.AcceptAlert()
	} else {
		err = wd.DismissAlert()
	}
	return err
}
