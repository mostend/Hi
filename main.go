package main

import (
	"Hi/client"
	_ "embed"

	"fmt"

	"github.com/getlantern/systray"
)

//go:embed icon/Hi.ico
var ic []byte

func main() {
	systray.Run(onReady, onExit)
}
func onReady() {
	systray.SetIcon(ic)
	systray.SetTitle("Hi")
	systray.SetTooltip("Hi")
	mXray := systray.AddMenuItem("主程序", "重新选择主程序")
	mConf := systray.AddMenuItem("配置", "重新选择配置文件")
	mQuit := systray.AddMenuItem("退出", "退出")
	//这里启动web server
	go func() {
		//router.Run()
	}()

	go func() {
		client.AppCmd.CheckCorePath()
		client.AppCmd.CheckConfPath()
		client.AppCmd.Start()
	}()
	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				go func() {
					client.AppCmd.Stop()
				}()
				systray.Quit()
				fmt.Println("Requesting quit")
				fmt.Println("Finished quitting")
				return

			case <-mXray.ClickedCh:
				go func() {
					client.AppCmd.Stop()
					client.AppCmd.ReCorePath()
					client.AppCmd.Start()
				}()
			case <-mConf.ClickedCh:
				//go func() {
				//	//打开web
				//	x.OpenBrowser(x.Browser())
				//}()
				go func() {
					client.AppCmd.Stop()
					client.AppCmd.ReConfPath()
					client.AppCmd.Start()
				}()
			}
		}
	}()
}

func onExit() {
	client.AppCmd.Stop()
	fmt.Println("exiting ...")
}
