package config


const appConfigFileName  ="app.properties"

var (
	//appconfig
	appConfig *AppConfig

)

type AppConfig struct {
	AppId string `json:"appId"`
	Cluster string `json:"cluster"`
	NamespaceName string `json:"namespaceName"`
	Ip string `json:"ip"`
	NextTryConnTime int64 `json:"-"`
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