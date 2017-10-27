package config


const appConfigFileName  ="app.properties"

var (
	//appconfig
	appConfig *AppConfig

)

type AppConfig struct {
	Zk string `json:"zk"`
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