package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gen2brain/dlgs"
	"github.com/go-cmd/cmd"
)

type App struct {
	Ctx      context.Context `json:"-"`
	Cmd      *cmd.Cmd        `json:"-"`
	CorePath string          `json:"CorePath"`
	ConfPath string          `json:"ConfPath"`
}

var AppCmd = NewApp(context.Background())

func NewApp(ctx context.Context) (a App) {
	a = App{
		Ctx: ctx,
	}
	CreateConfPath()
	userHomeDir, _ := os.UserHomeDir()

	if _, err := os.Stat(userHomeDir + AppConf); os.IsNotExist(err) {

		file, err := os.OpenFile(userHomeDir+AppConf, os.O_CREATE|os.O_WRONLY, 0777)
		defer file.Close()

		if err != nil {
			fmt.Println(err)
			dlgs.Error("Error", "打开app.json文件失败")
		}
		defaultAppConfig, _ := json.Marshal(a)
		file.Write(defaultAppConfig)
		return a
	}
	b, err := os.ReadFile(userHomeDir + AppConf)
	if err != nil {
		dlgs.Error("Error", "读取app.json失败")
	}
	a = App{}
	err = json.Unmarshal(b, &a)
	if err != nil {
		dlgs.Error("Error", "读取app.json失败")
	}
	return a
}

func (a *App) CheckCorePath() {
	if a.CorePath == "" {
		corePath, _, _ := dlgs.File("Core path", "", false)
		a.CorePath = corePath
	}
	a.SaveConf()
}
func (a *App) CheckConfPath() {
	if a.ConfPath == "" {
		confPath, _, _ := dlgs.File("config path", "", false)
		a.ConfPath = confPath
	}
	a.SaveConf()
}

func (a *App) ReCorePath() {
	corePath, ok, _ := dlgs.File("Core path", "", false)
	if ok {
		a.CorePath = corePath
		a.SaveConf()
	}
}

func (a *App) ReConfPath() {
	confPath, ok, _ := dlgs.File("config path", "", false)
	if ok {
		a.ConfPath = confPath
		a.SaveConf()
	}
}

func (a *App) SaveConf() {
	homeDir, _ := os.UserHomeDir()
	appFile := homeDir + AppConf
	file, err := os.OpenFile(appFile, os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		dlgs.Error("Error", "读取app.json失败")
	}
	newAppConf, _ := json.Marshal(a)
	file.Write(newAppConf)
}
