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
<<<<<<< HEAD
	mCore := systray.AddMenuItem("主程序", "重新选择主程序")
	mConf := systray.AddMenuItem("配置", "重新选择配置文件")
	mReStart := systray.AddMenuItem("重启", "重启")
	mQuit := systray.AddMenuItem("退出", "退出")
=======
	mCore := systray.AddMenuItem("Core", "Core")
	mConf := systray.AddMenuItem("Configuration", "Configuration")

	mShadowTls := systray.AddMenuItem("ShadowTLS", "ShadowTLS")
	mmShadowTlsConf := systray.AddMenuItem("ShadowTLSConf", "ShadowTLSConfiguration")

	mReStart := systray.AddMenuItem("Restart", "Restart")
	mQuit := systray.AddMenuItem("Quit", "Quit")
>>>>>>> 1667baa (shadowtls,bug has not been fixed)

	go func() {
		client.AppCmd.CheckPath()
		client.AppCmd.Start()
	}()

	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				go func() {
					client.AppCmd.Stop()
					client.AppCmd.DbClose()
				}()
				systray.Quit()
				fmt.Println("Requesting quit")
				fmt.Println("Finished quitting")
				return

			case <-mCore.ClickedCh:
				go func() {
					client.AppCmd.ReCore()
					client.AppCmd.ReStart()
				}()
			case <-mConf.ClickedCh:
				go func() {
					client.AppCmd.ReConf()
					client.AppCmd.ReStart()
				}()

			case <-mShadowTls.ClickedCh:
				go func() {
					client.AppCmd.ReShadowTlsCore()
					client.AppCmd.ReStart()
				}()
			case <-mmShadowTlsConf.ClickedCh:
				go func() {
					client.AppCmd.ReShadowTlsConf()
					client.AppCmd.ReStart()
				}()
			case <-mReStart.ClickedCh:
				go func() {
					client.AppCmd.ReStart()
				}()
			}

		}
	}()
}

func onExit() {
	client.AppCmd.Stop()
	client.AppCmd.DbClose()
	fmt.Println("exiting ...")
}
