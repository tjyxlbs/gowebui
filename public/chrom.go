package public

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func StartChromDriver() (selenium.WebDriver, error) {
	// 启动Chrome浏览器
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	chromeCaps := chrome.Capabilities{
		Args: []string{
			// "--headless", // 无头模式（可选）
		},
	}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515")
	if err != nil {
		return wd, err
	}
	return wd, nil
}
