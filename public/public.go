package public

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

const (
	SERVER  = "https://192.168.190.128:60443"
	ACCEPT  = "accept"
	DISMISS = "DISMISS"
)

type webDriverProcess func(wd selenium.WebDriver, args ...string) error

func SwitchToPage(url string) (selenium.WebDriver, error) {
	var err error
	wd, err := GetLogin()
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

func EleSendKeysByXpath(wd selenium.WebDriver, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("invalid args, expect args[0] key, args[1] value")
	}
	return EleSendKeys(wd, selenium.ByXPATH, args[0], args[1])
}

func EleSendKeys(wd selenium.WebDriver, args ...string) error {
	if len(args) != 3 {
		return fmt.Errorf("invalid args, expect args[0] by, args[1] key, args[2] value")
	}
	ele, err := wd.FindElement(args[0], args[1])
	if err != nil {
		return err
	}
	err = ele.Clear()
	if err != nil {
		return err
	}
	err = ele.SendKeys(args[2])
	if err != nil {
		return err
	}
	return nil
}

func EleClickByXpath(wd selenium.WebDriver, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid args, expect args[0] xpath")
	}
	return ClickBy(wd, selenium.ByXPATH, args[0])
}

func ClickBy(wd selenium.WebDriver, by, key string) error {
	// 下一步
	next, err := wd.FindElement(by, key)
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
