package jodz

import (
	"github.com/jodz/utils"
	"strings"
)

const appConfigFileName  ="app.properties"

var (
	//appconfig
	appConfig *AppConfig

)

func getAppConfig() *AppConfig {
	return appConfig
}

type AppConfig struct {
	Zk string `json:"zk"`
	zkIps []string
}

func (this *AppConfig)GetZkIps() []string {
	if utils.IsNotNil(this.zkIps){
		return this.zkIps
	}

	if utils.IsEmpty(this.Zk){
		return []string{}
	}

	this.zkIps=strings.Split(this.Zk,",")

	return this.zkIps
}


func init() {
	//init config
	initConfig()
}

func initConfig() {
	var err error
	//init config file
	appConfig,err = loadJsonConfig(appConfigFileName)

	if err!=nil{
		panic(err)
	}
}