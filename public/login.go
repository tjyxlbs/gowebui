package public

import (
	"github.com/tebeka/selenium"
)

var PublicWd selenium.WebDriver

func GetLogin() (selenium.WebDriver, error) {
	var err error
	if PublicWd != nil {
		return PublicWd, nil
	}
	PublicWd, err = StartChromDriver()
	if err != nil {
		return nil, err
	}
	// defer wd.Quit()
	PublicWd, err = Login(PublicWd)
	if err != nil {
		return nil, err
	}
	return PublicWd, nil
}

func Login(wd selenium.WebDriver) (selenium.WebDriver, error) {
	// 打开网页
	if err := wd.Get("https://ssl:K0a1@2023@192.168.190.128:60443/cert/ca/site_list"); err != nil {
		return wd, err
	}

	// 高级
	e0, err := wd.FindElement(selenium.ByXPATH, "/html/body/div/div[2]/button[3]")
	if err != nil {
		return wd, err
	}
	err = e0.Click()
	if err != nil {
		return wd, err
	}

	// 不安全继续
	e1, err := wd.FindElement(selenium.ByXPATH, "/html/body/div/div[3]/p[2]/a")
	if err != nil {
		return wd, err
	}
	err = e1.Click()
	if err != nil {
		return wd, err
	}

	// 超级管理员密码
	e2, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[4]/div[3]/div[2]/form/table/tbody/tr[2]/td[3]/input[2]")
	if err != nil {
		return wd, err
	}
	err = e2.SendKeys("12345678")
	if err != nil {
		return wd, err
	}

	// 提交
	e3, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[4]/div[3]/div[3]/span/img[1]")
	if err != nil {
		return wd, err
	}
	err = e3.Click()
	if err != nil {
		return wd, err
	}

	// 不修改密码
	err = wd.DismissAlert()
	if err != nil {
		return wd, err
	}
	return wd, nil
}
