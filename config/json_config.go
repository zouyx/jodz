package jodz

import (
	"io/ioutil"
	"encoding/json"
	"errors"
	"github.com/jodz/utils"
)

func loadJsonConfig(fileName string) (*AppConfig,error) {
	fs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil,errors.New("Fail to read config file:" + err.Error())
	}

	appConfig,loadErr:=createAppConfigWithJson(string(fs))

	if utils.IsNotNil(loadErr){
		return nil,errors.New("Load Json Config fail:" + loadErr.Error())
	}

	return appConfig,nil
}

func createAppConfigWithJson(str string) (*AppConfig,error) {
	appConfig:=&AppConfig{
	}
	err:=json.Unmarshal([]byte(str),appConfig)
	if utils.IsNotNil(err) {
		return nil,err
	}

	return appConfig,nil
}

