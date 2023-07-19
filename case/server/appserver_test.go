package server

import (
	"lbswebui/public"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/tebeka/selenium"
)

var appPage selenium.WebDriver

const (
	APP_LIST = public.SERVER + "/srv/trp/service"
)

func TestApp0(t *testing.T) {
	appPage, err := public.SwitchToPage(APP_LIST)
	assert.Equal(t, err, nil)
	err = deleteAllApp(appPage, APP_LIST)
	assert.Equal(t, err, nil)
}

func TestApp1(t *testing.T) {
	appPage, err := public.SwitchToPage(APP_LIST)
	assert.Equal(t, err, nil)
	err = appAdd(appPage)
	assert.Equal(t, err, nil)
	// 服务名
	err = public.TextEle(appPage, "service_name", "rsa-111")
	assert.Equal(t, err, nil)
	// 端口
	err = public.TextEle(appPage, "local_port", "20444")
	assert.Equal(t, err, nil)
	// 第一步，next
	err = next1(appPage)
	assert.Equal(t, err, nil)
	// 真实服务ip realserver_ip[]
	err = public.TextEle(appPage, "realserver_ip[]", "192.168.190.200")
	assert.Equal(t, err, nil)
	// 真实服务port realserver_port[]
	err = public.TextEle(appPage, "realserver_port[]", "9191")
	assert.Equal(t, err, nil)
	// 第二步，next
	err = next2(appPage)
	assert.Equal(t, err, nil)
	// 选择a.hello.com证书
	// ByCSSSelector
	rsaSites, err := appPage.FindElements(selenium.ByCSSSelector, "option[keytype=\"RSA\"]")
	assert.Equal(t, err, nil)
	assert.NotEqual(t, len(rsaSites), 0)
	bFind := false
	for _, site := range rsaSites {
		s, err := site.Text()
		if err != nil {
			continue
		}
		if strings.Contains(s, "test.com") {
			err = site.Click()
			bFind = true
			assert.Equal(t, err, nil)
		}
	}
	assert.Equal(t, bFind, true)
	// 保存
	err = saveApp(appPage)
	assert.Equal(t, err, nil)
}

func TestApp2(t *testing.T) {
	appPage, err := public.SwitchToPage(APP_LIST)
	assert.Equal(t, err, nil)
	err = appAdd(appPage)
	assert.Equal(t, err, nil)
	// 服务名
	err = public.TextEle(appPage, "service_name", "rsa-222")
	assert.Equal(t, err, nil)
	// 端口
	err = public.TextEle(appPage, "local_port", "20445")
	assert.Equal(t, err, nil)
	// 第一步，next
	err = next1(appPage)
	assert.Equal(t, err, nil)
	// 真实服务ip realserver_ip[]
	err = public.TextEle(appPage, "realserver_ip[]", "192.168.190.200")
	assert.Equal(t, err, nil)
	// 真实服务port realserver_port[]
	err = public.TextEle(appPage, "realserver_port[]", "9192")
	assert.Equal(t, err, nil)
	// 第二步，next
	err = next2(appPage)
	assert.Equal(t, err, nil)
	// 选择第二个证书
	rsaSites, err := appPage.FindElements(selenium.ByCSSSelector, "option[keytype=\"RSA\"]")
	assert.Equal(t, err, nil)
	assert.NotEqual(t, len(rsaSites), 0)
	bFind := false
	for _, site := range rsaSites {
		s, err := site.Text()
		if err != nil {
			continue
		}
		if strings.Contains(s, "lbs-rsa-cs.com") {
			err = site.Click()
			bFind = true
			assert.Equal(t, err, nil)
		}
	}
	assert.Equal(t, bFind, true)
	// 保存
	err = saveApp(appPage)
	assert.Equal(t, err, nil)
}

func deleteAllApp(wd selenium.WebDriver, url string) error {
	// 获取app列表
	appLst, err := wd.FindElements(selenium.ByName, "data_id[]")
	if err != nil {
		return err
	}
	for _, app := range appLst {
		err = app.Click()
		if err != nil {
			continue
		}
	}
	// 删除
	err = public.EleClickByXpath(wd, "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/form/table[2]/tbody/tr/td/a[1]/span")
	if err != nil {
		return err
	}
	// 等待弹窗
	return public.WaitAlert(wd, public.ACCEPT, "")
}

func appAdd(wd selenium.WebDriver) error {
	add, err := wd.FindElement(selenium.ByXPATH, "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/form/table[2]/tbody/tr/td/a[2]/span")
	if err != nil {
		return err
	}
	err = add.Click()
	if err != nil {
		return err
	}
	// 下一步
	next, err := wd.FindElement(selenium.ByXPATH, "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/a[1]/span")
	if err != nil {
		return err
	}
	err = next.Click()
	if err != nil {
		return err
	}
	return nil
}

func next1(wd selenium.WebDriver) error {
	return public.EleClickByXpath(wd, "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/table/tbody/tr/td/a[1]/span")
}

func next2(wd selenium.WebDriver) error {
	return public.EleClickByXpath(wd, "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table[2]/tbody/tr/td/a[2]/span")
}

func saveApp(wd selenium.WebDriver) error {
	return public.EleClickByXpath(wd, "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/a[1]/span")
}
