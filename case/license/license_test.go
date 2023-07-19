package server

import (
	"lbswebui/public"
	"os"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/tebeka/selenium"
)

const (
	LICENSE_EFCA_URL = "http://gef.koal.com:8080/#/license/index"
	LICENSE_AISG_URL = public.SERVER + "/sys/license/license"
	DOWN_LOAD_PATH   = "C:\\Users\\Lenovo\\Downloads"
	SEQ              = "44164D56-6B73-C47B-E5C1-32AA63FA900C"
)

func TestApp0(t *testing.T) {
	wdLicense, err := public.StartChromDriver()
	assert.Equal(t, err, nil)

	// 进入效能工厂
	err = wdLicense.Get(LICENSE_EFCA_URL)
	assert.Equal(t, err, nil)

	// 输入序列号
	seq := "/html/body/div[1]/div/section/section/main/section/main/div/div/div[1]/div/div/form/div[1]/div[2]/div[1]/div/div/input"
	err = public.WaitAndDo(wdLicense, public.EleSendKeysByXpath, seq, SEQ)
	assert.Equal(t, err, nil)

	// 选择虚拟机
	vm := "/html/body/div/div/section/section/main/section/main/div/div/div[1]/div/div/form/div[4]/div[2]/div/div/label[1]/span[2]"
	err = public.WaitAndDo(wdLicense, public.EleClickByXpath, vm)
	assert.Equal(t, err, nil)

	// 点击授权项
	cfg := "/html/body/div[1]/div/section/section/main/section/main/div/div/div[1]/div/div/form/div[6]/div[2]/div/div/div/div/input"
	err = public.WaitAndDo(wdLicense, public.EleClickByXpath, cfg)
	assert.Equal(t, err, nil)

	// 选择aisg
	aisg := "/html/body/div[2]/div/div/div/div/ul[2]/li[3]/span"
	err = public.WaitAndDo(wdLicense, public.EleClickByXpath, aisg)
	assert.Equal(t, err, nil)

	// 使用配置模版
	use := "/html/body/div[1]/div/section/section/main/section/main/div/div/div[1]/div/div/form/div[6]/div[2]/div/button[1]/span"
	err = public.WaitAndDo(wdLicense, public.EleClickByXpath, use)
	assert.Equal(t, err, nil)

	// 下载license
	download := "/html/body/div[1]/div/section/section/main/section/main/div/div/div[1]/div/div/form/div[9]/div[2]/div/button/span"
	timeBeforDown := time.Now()
	err = public.WaitAndDo(wdLicense, public.EleClickByXpath, download)
	assert.Equal(t, err, nil)

	// 等待下载文件4s
	lic, err := public.WaitDownloadFile(wdLicense, DOWN_LOAD_PATH, "*.lic", timeBeforDown)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, lic, "")

	// 关闭效能工厂浏览器
	err = wdLicense.Quit()
	assert.Equal(t, err, nil)

	// 进入上传license界面
	aisgWeb, err := public.GetLogin()
	assert.Equal(t, err, nil)
	assert.NotEqual(t, aisgWeb, nil)
	err = aisgWeb.Get(LICENSE_AISG_URL)
	assert.Equal(t, err, nil)

	// 导入lic
	upload := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/table[2]/tbody/tr/td/form/table/tbody/tr/td[1]/div/div/input[2]"
	err = public.WaitAndDo(aisgWeb, public.EleSendKeysByXpath, upload, lic)
	assert.Equal(t, err, nil)

	// 点击确认导入
	err = public.WaitAlert(aisgWeb, public.ACCEPT, "导入新的许可证文件？")
	assert.Equal(t, err, nil)

	// 判断是否导入成功
	version := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/table[1]/tbody/tr[2]/td[1]"
	ele, err := aisgWeb.FindElement(selenium.ByXPATH, version)
	assert.Equal(t, err, nil)
	displayed, err := ele.IsDisplayed()
	assert.Equal(t, err, nil)
	assert.Equal(t, displayed, true)

	// 删除lic
	err = os.Remove(lic)
	assert.Equal(t, err, nil)
}
