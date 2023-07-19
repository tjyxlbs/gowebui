package cms

import (
	"fmt"
	"lbswebui/public"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
)

type certData struct {
	Cert     string `json:"cert"`
	Key      string `json:"key"`
	Password string `json:"password"`
}

var certMap map[string]certData

const (
	SITE_LIST    = SERVER + "/cert/ca/site_list"
	BACK_UPLOAD  = "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/table/tbody/tr/td/a[2]/span"
	BACK_CONTENT = "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/table[2]/tbody/tr/td/form/table/tbody/tr/td/a[2]/span"
)

func init() {
	certMap = make(map[string]certData)
	certMap["c1"] = certData{
		Cert:     "D:\\File\\work\\task\\GAD测试\\验证扩展用法\\server2\\cert_rsa_2048_lbs-cs-2.com_4688d7aa-bd26-489c-91a1-c8ad21af8182.pem",
		Key:      "D:\\File\\work\\task\\GAD测试\\验证扩展用法\\server2\\cert_rsa_2048_lbs-cs-2.com_4688d7aa-bd26-489c-91a1-c8ad21af8182.key",
		Password: "",
	}

	sm2Dir := "D:\\File\\work\\task\\冒烟测试\\sm2-gad-1\\server"
	certMap["c2-enc"] = certData{
		Cert:     fmt.Sprintf("%s/%s", sm2Dir, "gad_test1.com-site-enc.pem"),
		Key:      fmt.Sprintf("%s/%s", sm2Dir, "gad_test1.com-site-enc.key"),
		Password: "",
	}
	certMap["c2-sig"] = certData{
		Cert:     fmt.Sprintf("%s/%s", sm2Dir, "gad_test1.com-site-sig.pem"),
		Key:      fmt.Sprintf("%s/%s", sm2Dir, "gad_test1.com-site-sig.key"),
		Password: "",
	}
	pfxDir := "D:\\File\\work"
	certMap["c3-pfx"] = certData{
		Cert:     fmt.Sprintf("%s/%s", pfxDir, "lbs1.pfx"),
		Key:      "",
		Password: "lbs@123.com",
	}
}

func TestC0(t *testing.T) {
	wd, err := public.SwitchToPage(SITE_LIST)
	assert.Equal(t, err, nil)
	err = DeleteAllCert(wd, SITE_LIST)
	assert.Equal(t, err, nil)
	ele, err := wd.FindElement(selenium.ByXPATH, "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/table/tbody/tr/td/form/table/tbody/tr/td[1]/div/a/span")
	if err != nil {
		// 证书被其他服务引用，弹窗显示
		err = WaitAlert(wd, public.ACCEPT, "")
		assert.Equal(t, err, nil)
	} else {
		ok, err := ele.IsDisplayed()
		if err != nil || !ok {
			// 证书被其他服务引用，弹窗显示
			err = WaitAlert(wd, public.ACCEPT, "")
			assert.Equal(t, err, nil)
		}
	}
}

func TestC1(t *testing.T) {
	sitePage, err := public.SwitchToPage(SITE_LIST)
	assert.Equal(t, err, nil)
	err = gotoUpload(sitePage) // 进入证书导入界面
	assert.Equal(t, err, nil)

	err = uploadCert(sitePage, certMap["c1"]) // 导入证书
	assert.Equal(t, err, nil)

	err = uploadClick(sitePage) // 点击上传
	assert.Equal(t, err, nil)

	err = back(sitePage, BACK_UPLOAD) // 点击返回
	assert.Equal(t, err, nil)

	_, err = sitePage.FindElement(selenium.ByLinkText, "lbs-cs-2.com")
	assert.Equal(t, err, nil)
}

func TestC2(t *testing.T) {
	sitePage, err := public.SwitchToPage(SITE_LIST)
	assert.Equal(t, err, nil)
	err = gotoUpload(sitePage) // 进入证书导入界面
	assert.Equal(t, err, nil)

	err = uploadCert(sitePage, certMap["c2-sig"]) // 导入签名证书
	assert.Equal(t, err, nil)

	err = uploadClick(sitePage) // 点击上传
	assert.Equal(t, err, nil)

	err = uploadCert(sitePage, certMap["c2-enc"]) // 导入加密证书
	assert.Equal(t, err, nil)

	err = uploadClick(sitePage) // 点击上传
	assert.Equal(t, err, nil)

	err = back(sitePage, BACK_UPLOAD) // 点击返回
	assert.Equal(t, err, nil)

	enc, err := sitePage.FindElement(selenium.ByLinkText, "gad_test1.com[加密证书]")
	assert.Equal(t, err, nil)
	err = enc.Click() // 进入详情界面
	assert.Equal(t, err, nil)
	err = back(sitePage, BACK_CONTENT) // 返回证书列表
	assert.Equal(t, err, nil)

	sig, err := sitePage.FindElement(selenium.ByLinkText, "gad_test1.com[签名证书]")
	assert.Equal(t, err, nil)
	err = sig.Click() // 进入详情界面
	assert.Equal(t, err, nil)
	err = back(sitePage, BACK_CONTENT) // 返回证书列表
	assert.Equal(t, err, nil)
}
func TestC3(t *testing.T) {
	sitePage, err := public.SwitchToPage(SITE_LIST)
	assert.Equal(t, err, nil)
	err = gotoUpload(sitePage) // 进入证书导入界面
	assert.Equal(t, err, nil)

	err = uploadCert(sitePage, certMap["c3-pfx"]) // 导入pfx
	assert.Equal(t, err, nil)

	err = uploadClick(sitePage) // 点击上传
	assert.Equal(t, err, nil)

	err = back(sitePage, BACK_UPLOAD) // 点击返回
	assert.Equal(t, err, nil)

	_, err = sitePage.FindElement(selenium.ByLinkText, "李本帅")
	assert.Equal(t, err, nil)
}

// 异常测试：导入证书链
func TestC4(t *testing.T) {
	sitePage, err := public.SwitchToPage(SITE_LIST)
	assert.Equal(t, err, nil)
	err = gotoUpload(sitePage) // 进入证书导入界面
	assert.Equal(t, err, nil)

	err = uploadCert(sitePage, certData{
		Cert: "D:\\File\\work\\task\\冒烟测试\\rsa\\ca\\sadg_rsa_local.pem", // 导入证书链
	}) // 导入pfx
	assert.Equal(t, err, nil)

	err = uploadClick(sitePage) // 点击上传
	assert.Equal(t, err, nil)

	// 错误弹框
	err = WaitAlert(sitePage, public.DISMISS, "不能导入证书链")
	assert.Equal(t, err, nil)

	err = back(sitePage, BACK_UPLOAD) // 点击返回
	assert.Equal(t, err, nil)
}

// 导入der格式证书，没有秘钥导入失败
func TestC5(t *testing.T) {
	sitePage, err := public.SwitchToPage(SITE_LIST)
	assert.Equal(t, err, nil)
	err = gotoUpload(sitePage) // 进入证书导入界面
	assert.Equal(t, err, nil)

	err = uploadCert(sitePage, certData{
		Cert: "D:\\File\\work\\task\\冒烟测试\\rsa\\server\\aaa.der", // 导入der证书
	}) // 导入pfx
	assert.Equal(t, err, nil)

	err = uploadClick(sitePage) // 点击上传
	assert.Equal(t, err, nil)

	// 错误弹框
	err = WaitAlert(sitePage, public.DISMISS, "导入证书失败")
	assert.Equal(t, err, nil)

	err = back(sitePage, BACK_UPLOAD) // 点击返回
	assert.Equal(t, err, nil)
}

func TestC6(t *testing.T) {
	sitePage, err := public.SwitchToPage(SITE_LIST)
	assert.Equal(t, err, nil)
	err = gotoUpload(sitePage) // 进入证书导入界面
	assert.Equal(t, err, nil)

	err = uploadCert(sitePage, certData{
		Cert: "D:\\File\\work\\task\\冒烟测试\\rsa\\server\\aaa.der",      // 导入der证书
		Key:  "D:\\File\\work\\task\\冒烟测试\\rsa\\server\\test.com.key", // 导入私钥
	})
	assert.Equal(t, err, nil)

	err = uploadClick(sitePage) // 点击上传
	assert.Equal(t, err, nil)

	err = back(sitePage, BACK_UPLOAD) // 点击返回
	assert.Equal(t, err, nil)

	_, err = sitePage.FindElement(selenium.ByLinkText, "test.com")
	assert.Equal(t, err, nil)
}

func TestC7(t *testing.T) {
	sitePage, err := public.SwitchToPage(SITE_LIST)
	assert.Equal(t, err, nil)
	err = gotoUpload(sitePage) // 进入证书导入界面
	assert.Equal(t, err, nil)

	err = uploadCert(sitePage, certData{
		Cert: "D:\\File\\work\\task\\冒烟测试\\1024-RSA\\a.hello.com-RSA-1024.pem", // 导入证书
		Key:  "D:\\File\\work\\task\\冒烟测试\\1024-RSA\\a.hello.com-RSA-1024.key", // 导入私钥
	})
	assert.Equal(t, err, nil)

	err = uploadClick(sitePage) // 点击上传
	assert.Equal(t, err, nil)

	err = back(sitePage, BACK_UPLOAD) // 点击返回
	assert.Equal(t, err, nil)
}

// D:\File\work\task\GAD测试\验证扩展用法\ca
func TestC8(t *testing.T) {
	sitePage, err := public.SwitchToPage(SITE_LIST)
	assert.Equal(t, err, nil)
	err = gotoUpload(sitePage) // 进入证书导入界面
	assert.Equal(t, err, nil)

	err = uploadCert(sitePage, certData{
		Cert: "D:\\File\\work\\task\\GAD测试\\验证扩展用法\\server\\cert_rsa_2048_lbs-rsa-cs.com_d55e9c5e-028c-44a7-9e42-043551e4b7ab.pem", // 导入证书
		Key:  "D:\\File\\work\\task\\GAD测试\\验证扩展用法\\server\\cert_rsa_2048_lbs-rsa-cs.com_d55e9c5e-028c-44a7-9e42-043551e4b7ab.key", // 导入私钥
	})
	assert.Equal(t, err, nil)

	err = uploadClick(sitePage) // 点击上传
	assert.Equal(t, err, nil)

	err = back(sitePage, BACK_UPLOAD) // 点击返回
	assert.Equal(t, err, nil)
}

func uploadCert(wd selenium.WebDriver, certObj certData) error {

	cert := certObj.Cert
	key := certObj.Key
	password := certObj.Password
	if cert != "" {
		// 导入证书
		upCert, err := wd.FindElement(selenium.ByID, "certfile")
		if err != nil {
			return err
		}
		err = upCert.SendKeys(cert)
		if err != nil {
			return err
		}
	}

	if key != "" {
		// 导入私钥
		upKey, err := wd.FindElement(selenium.ByID, "key")
		if err != nil {
			return err
		}
		err = upKey.SendKeys(key)
		if err != nil {
			return err
		}
	}

	//passwd
	if password != "" {
		// 导入密码
		upPassword, err := wd.FindElement(selenium.ByID, "passwd")
		if err != nil {
			return err
		}
		err = upPassword.SendKeys(password)
		if err != nil {
			return err
		}
	}
	return nil
}

func gotoUpload(wd selenium.WebDriver) error {
	up, err := wd.FindElement(selenium.ByID, "btn_upload")
	if err != nil {
		return err
	}
	err = up.Click()
	if err != nil {
		return err
	}
	return nil
}
func back(wd selenium.WebDriver, xpath string) error {
	back, err := wd.FindElement(selenium.ByXPATH, xpath)
	if err != nil {
		return err
	}
	err = back.Click()
	if err != nil {
		return err
	}
	return nil
}
func uploadClick(wd selenium.WebDriver) error {
	// 上传
	up, err := wd.FindElement(selenium.ByXPATH, "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/table/tbody/tr/td/a[1]/span")
	if err != nil {
		return err
	}
	err = up.Click()
	if err != nil {
		return err
	}
	return nil
}
