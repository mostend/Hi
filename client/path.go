package client

import "os"

var AppPath = "/.Hi/"
var AppConf = "/.Hi/app.json"
var RunConf = "/.Hi/config/"
var CustomerPath = "/.Hi/Customer/"

//var ConfFiles = []string{
//	"00_log.json",
//	"01_api.json",
//	"02_dns.json",
//	"03_routing.json",
//	"04_policy.json",
//	"05_inbounds.json",
//	"06_outbounds.json",
//	"07_transport.json",
//	"08_stats.json",
//	"09_reverse.json",
//}

func CreateConfPath() {

	userHomeDir, _ := os.UserHomeDir()
	if _, err := os.Stat(userHomeDir + AppPath); os.IsNotExist(err) {
		os.MkdirAll(userHomeDir+AppPath, 0777)
	}
	if _, err := os.Stat(userHomeDir + RunConf); os.IsNotExist(err) {
		os.MkdirAll(userHomeDir+RunConf, 0777)
	}
	if _, err := os.Stat(userHomeDir + CustomerPath); os.IsNotExist(err) {
		os.MkdirAll(userHomeDir+CustomerPath, 0777)
	}
}
