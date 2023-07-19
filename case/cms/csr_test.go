package cms

import (
	"lbswebui/public"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/tebeka/selenium"
)

const (
	CSR_PAGE = SERVER + "/cert/ca/csr"
)

var csrPage selenium.WebDriver

func testConfigSet(t *testing.T, wd selenium.WebDriver,
	kvPair map[string]string, by string) {
	for k, v := range kvPair {
		id, err := wd.FindElement(by, k)
		assert.Equal(t, err, nil)
		err = id.SendKeys(v)
		assert.Equal(t, err, nil)
	}
}

func TestCsr1(t *testing.T) {
	wd, err := public.SwitchToPage(CSR_PAGE)
	assert.Equal(t, err, nil)
	idInfo := make(map[string]string)
	// 站点名称
	idInfo["commonName"] = "TestCsr1.com"
	// 公司
	idInfo["organizationName"] = "case1Compony"
	// 部门 organizationalUnitName
	idInfo["organizationalUnitName"] = "gw1"
	// 电子邮件 emailAddress
	idInfo["emailAddress"] = "csrtest1@koal.com"
	// 私钥密码 password
	idInfo["password"] = "123123123"
	testConfigSet(t, wd, idInfo, selenium.ByID)
	// 密钥类型 RSA
	key, err := wd.FindElement(selenium.ByXPATH, "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/form/table/tbody/tr[8]/td[2]/select/option[2]")
	assert.Equal(t, err, nil)
	err = key.Click()
	assert.Equal(t, err, nil)
	// 密钥长度
	keyLength, err := wd.FindElement(selenium.ByXPATH, "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/form/table/tbody/tr[9]/td[2]/select/option[3]")
	assert.Equal(t, err, nil)
	err = keyLength.Click()
	assert.Equal(t, err, nil)
	// 提交 btnGenerateReq
	submit, err := wd.FindElement(selenium.ByID, "btnGenerateReq")
	assert.Equal(t, err, nil)
	err = submit.Click()
	assert.Equal(t, err, nil)
	// 返回
	err = wd.Back()
	assert.Equal(t, err, nil)
}

func TestCsr2(t *testing.T) {
	wd, err := public.SwitchToPage(CSR_PAGE)
	assert.Equal(t, err, nil)
	idInfo := make(map[string]string)
	// 站点名称
	idInfo["commonName"] = "TestCsr2.com"
	// 公司
	idInfo["organizationName"] = "case2Compony"
	// 部门 organizationalUnitName
	idInfo["organizationalUnitName"] = "gw2"
	// 电子邮件 emailAddress
	idInfo["emailAddress"] = "csrtest2@koal.com"
	// 私钥密码 password
	idInfo["password"] = "123123123"
	testConfigSet(t, wd, idInfo, selenium.ByID)
	// 密钥类型 SM2
	key, err := wd.FindElement(selenium.ByXPATH, "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/form/table/tbody/tr[8]/td[2]/select/option[1]")
	assert.Equal(t, err, nil)
	err = key.Click()
	assert.Equal(t, err, nil)
	// 提交 btnGenerateReq
	submit, err := wd.FindElement(selenium.ByID, "btnGenerateReq")
	assert.Equal(t, err, nil)
	err = submit.Click()
	assert.Equal(t, err, nil)
	// 返回
	err = wd.Back()
	assert.Equal(t, err, nil)
}
