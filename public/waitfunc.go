package public

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/tebeka/selenium"
)

func WaitAndDo(wd selenium.WebDriver, f webDriverProcess, args ...string) (err error) {
	var IsSet = func(wd selenium.WebDriver) (bool, error) {
		err := f(wd, args...)
		if err != nil {
			return false, nil
		}
		return true, nil
	}
	// 超时时间和间隔时间
	err = wd.WaitWithTimeoutAndInterval(IsSet, 5*time.Second, 500*time.Millisecond)
	if err != nil {
		return err
	}
	return nil
}

func WaitAndSendKeys(wd selenium.WebDriver, by, key, value string) (err error) {
	var IsSet = func(wd selenium.WebDriver) (bool, error) {
		err := EleSendKeys(wd, by, key, value)
		if err != nil {
			return false, nil
		}
		return true, nil
	}
	// 超时时间和间隔时间
	err = wd.WaitWithTimeoutAndInterval(IsSet, 10*time.Second, 500*time.Millisecond)
	if err != nil {
		return err
	}
	return nil
}

func WaitDownloadFile(wd selenium.WebDriver, folderPath, pattern string, beforDown time.Time) (downLoadFile string, err error) {
	var IsDownload = func(wd selenium.WebDriver) (bool, error) {
		downLoadFile, err = waitDownLoadFile(folderPath, pattern, beforDown)
		if err != nil {
			return false, nil
		}
		return true, nil
	}
	// 超时时间和间隔时间
	err = wd.WaitWithTimeoutAndInterval(IsDownload, 10*time.Second, 500*time.Millisecond)
	if err != nil {
		return
	}
	return
}

func waitDownLoadFile(folderPath, pattern string, beforDown time.Time) (filePath string, err error) {
	files, err := filepath.Glob(fmt.Sprintf("%s\\%s", folderPath, pattern))
	if err != nil {
		return "", err
	}
	if len(files) < 1 {
		return "", fmt.Errorf("no file")
	}
	lastFile := files[0]
	info, err := os.Stat(lastFile)
	if err != nil {
		return "", err
	}
	lastTime := info.ModTime()
	for i := 1; i < len(files); i++ {
		file := files[i]
		info, err := os.Stat(file)
		if err != nil {
			continue
		}
		if info.ModTime().Unix() > lastTime.Unix() {
			lastTime = info.ModTime()
			lastFile = file
		}
	}
	if lastTime.Unix() > beforDown.Unix() {
		return lastFile, nil
	}
	return "", fmt.Errorf("no file")
}
