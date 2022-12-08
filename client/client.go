package client

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/gen2brain/dlgs"

	"github.com/xujiajun/nutsdb"

	"github.com/go-cmd/cmd"
)

type App struct {
	Ctx  context.Context
	Cmd  *cmd.Cmd
	Db   *nutsdb.DB
	Core string
	Conf string
}

var AppCmd = NewApp(context.Background())

func NewApp(ctx context.Context) (a *App) {
	a = &App{
		Ctx: ctx,
	}
	userHomeDir, _ := os.UserHomeDir()
	dbPath := userHomeDir + AppConf
	opt := nutsdb.DefaultOptions
	opt.Dir = dbPath
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	a.Db = db
	return a
}

func (a *App) CheckPath() {
	if err := a.Db.Update(
		func(tx *nutsdb.Tx) error {
			coreKey := []byte(Core)
			confKey := []byte(Conf)
			bucket := BucketOfPath
			if e, err := tx.Get(bucket, coreKey); err != nil {
				corePath, _, _ := dlgs.File("Core path", "", false)
				a.Core = corePath
				tx.Put(bucket, coreKey, []byte(a.Core), 0)
			} else {
				a.Core = string(e.Value)
			}
			if e, err := tx.Get(bucket, confKey); err != nil {
				confPath, _, _ := dlgs.File("Config path", "", false)
				a.Conf = confPath
				tx.Put(bucket, confKey, []byte(a.Conf), 0)
			} else {
				a.Conf = string(e.Value)
			}
			return nil
		}); err != nil {
		dlgs.Error("Error", "读取配置失败")
	}
}

func (a *App) ReCore() {
	corePath, ok, _ := dlgs.File("Core path", "", false)
	if ok {
		a.Core = corePath
		a.SaveConf()
	}
}

func (a *App) ReConf() {
	confPath, ok, _ := dlgs.File("config path", "", false)
	if ok {
		a.Conf = confPath
		a.SaveConf()
	}
}

func (a *App) SaveConf() {
	if err := a.Db.Update(
		func(tx *nutsdb.Tx) error {
			coreKey := []byte(Core)
			confKey := []byte(Conf)
			bucket := BucketOfPath
			if err := tx.Put(bucket, coreKey, []byte(a.Core), 0); err != nil {
				return err
			}
			if err := tx.Put(bucket, confKey, []byte(a.Conf), 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		dlgs.Error("Error", err.Error())
	}
}

func (a *App) GetRunCmd() []string {

	if strings.Contains(a.Core, "sing-box") {
		return []string{"run", "-c", a.Conf}
	}
	return []string{"-c", a.Conf}
}

func (a *App) DbClose() {
	a.Db.Close()
}
