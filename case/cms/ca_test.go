package cms

import (
	"lbswebui/public"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/tebeka/selenium"
)

var caPage selenium.WebDriver

const (
	CA_LIST = SERVER + "/cert/ca/ca_list"
)

func TestCa0(t *testing.T) {
	caPage, err := public.SwitchToPage(CA_LIST)
	assert.Equal(t, err, nil)
	err = DeleteAllCert(caPage, CA_LIST)
	assert.Equal(t, err, nil)
	ele, err := caPage.FindElement(selenium.ByXPATH, "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/table/tbody/tr/td/form/table/tbody/tr/td[2]/a/span")
	if err != nil {
		// 证书被其他服务引用，弹窗显示
		err = WaitAlert(caPage, public.ACCEPT, "")
		assert.Equal(t, err, nil)
	} else {
		ok, err := ele.IsDisplayed()
		if err != nil || !ok {
			// 证书被其他服务引用，弹窗显示
			err = WaitAlert(caPage, public.ACCEPT, "")
			assert.Equal(t, err, nil)
		}
	}
}

// 导入证书失败 上级证书是否导入？
func TestCa1(t *testing.T) {
	var err error
	caPage, err := public.SwitchToPage(CA_LIST)
	assert.Equal(t, err, nil)
	err = caImport(caPage, "D:\\File\\work\\task\\GAD测试\\验证扩展用法\\ca\\ca_rsa_lbs-gad-local.com_3077a71e-0b5e-40fe-8662-6d7999d73f00.pem")
	assert.Equal(t, err, nil)
	// 关闭弹窗
	err = WaitAlert(caPage, public.ACCEPT, "导入证书失败, 上级证书是否导入？")
	assert.Equal(t, err, nil)
}

// 导入root->local
func TestCa2(t *testing.T) {
	var err error
	caPage, err := public.SwitchToPage(CA_LIST)
	assert.Equal(t, err, nil)

	err = caImport(caPage, "D:\\File\\work\\task\\GAD测试\\验证扩展用法\\ca\\ca_rsa_lbs-gad-root.com_061930b3-f90c-46c0-bc4e-c5289d24388b.pem")
	assert.Equal(t, err, nil)
	// 关闭成功弹窗
	err = WaitAlert(caPage, public.ACCEPT, "导入成功")
	assert.Equal(t, err, nil)

	err = caImport(caPage, "D:\\File\\work\\task\\GAD测试\\验证扩展用法\\ca\\ca_rsa_lbs-gad-local.com_3077a71e-0b5e-40fe-8662-6d7999d73f00.pem")
	assert.Equal(t, err, nil)
	// 关闭成功弹窗
	err = WaitAlert(caPage, public.ACCEPT, "导入成功")
	assert.Equal(t, err, nil)
}

// 导入sm2证书链
func TestCa3(t *testing.T) {
	var err error
	caPage, err := public.SwitchToPage(CA_LIST)
	assert.Equal(t, err, nil)
	// 导入root证书
	err = caImport(caPage,
		"D:\\File\\work\\task\\冒烟测试\\sm2-gad-1\\ca\\gad_test_sm2_root.pem")
	assert.Equal(t, err, nil)
	// 关闭成功弹窗
	err = WaitAlert(caPage, public.ACCEPT, "导入成功")
	assert.Equal(t, err, nil)

	// 导入local证书
	err = caImport(caPage,
		"D:\\File\\work\\task\\冒烟测试\\sm2-gad-1\\ca\\gad_test_sm2_local.pem")
	assert.Equal(t, err, nil)
	// 关闭成功弹窗
	err = WaitAlert(caPage, public.ACCEPT, "导入成功")
	assert.Equal(t, err, nil)
}

func TestCa4(t *testing.T) {
	var err error
	caPage, err := public.SwitchToPage(CA_LIST)
	assert.Equal(t, err, nil)
	// 导入root证书
	err = caImport(caPage,
		"D:\\File\\work\\task\\GAD测试\\验证扩展用法\\ca\\ca_rsa_lbs-gad-root.com_061930b3-f90c-46c0-bc4e-c5289d24388b.pem")
	assert.Equal(t, err, nil)
	// 关闭成功弹窗
	err = WaitAlert(caPage, public.ACCEPT, "导入成功")
	assert.Equal(t, err, nil)

	// 导入local证书
	err = caImport(caPage,
		"D:\\File\\work\\task\\GAD测试\\验证扩展用法\\ca\\ca_rsa_lbs-gad-local.com_3077a71e-0b5e-40fe-8662-6d7999d73f00.pem")
	assert.Equal(t, err, nil)
	// 关闭成功弹窗
	err = WaitAlert(caPage, public.ACCEPT, "导入成功")
	assert.Equal(t, err, nil)
}

func TestCa5(t *testing.T) {
	var err error
	caPage, err := public.SwitchToPage(CA_LIST)
	assert.Equal(t, err, nil)
	// 导入root证书
	err = caImport(caPage,
		"D:\\File\\work\\task\\冒烟测试\\rsa\\ca\\sadg_rsa_root.pem")
	assert.Equal(t, err, nil)
	// 关闭成功弹窗
	err = WaitAlert(caPage, public.ACCEPT, "导入成功")
	assert.Equal(t, err, nil)

	// 导入local证书
	err = caImport(caPage,
		"D:\\File\\work\\task\\冒烟测试\\rsa\\ca\\sadg_rsa_local.pem")
	assert.Equal(t, err, nil)
	// 关闭成功弹窗
	err = WaitAlert(caPage, public.ACCEPT, "导入成功")
	assert.Equal(t, err, nil)
}

func caImport(wd selenium.WebDriver, cert string) error {
	certEle, err := wd.FindElement(selenium.ByID, "certfile")
	if err != nil {
		return err
	}
	err = certEle.SendKeys(cert)
	if err != nil {
		return err
	}
	return nil
}
