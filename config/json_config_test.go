package config

import (
	"testing"
	"container/list"
	"github.com/zouyx/jodz/test"
)

func TestLoadJsonConfig(t *testing.T) {
	config,err:=loadJsonConfig(appConfigFileName)
	t.Log(config)

	test.Nil(t,err)
	test.NotNil(t,config)
	test.Equal(t,"127.0.0.1:2180",config.Zk)

}

func TestLoadJsonConfigWrongFile(t *testing.T) {
	config,err:=loadJsonConfig("")
	test.NotNil(t,err)
	test.Nil(t,config)

	test.StartWith(t,"Fail to read config file",err.Error())
}

func TestLoadJsonConfigWrongType(t *testing.T) {
	config,err:=loadJsonConfig("app_config.go")
	test.NotNil(t,err)
	test.Nil(t,config)

	test.StartWith(t,"Load Json Config fail",err.Error())
}

func TestCreateAppConfigWithJson(t *testing.T) {
	jsonStr:=`{
    "zk": "test",
    "cluster": "dev",
    "namespaceName": "application",
    "ip": "localhost:8888",
    "releaseKey": ""
	}`
	config,err:=createAppConfigWithJson(jsonStr)
	t.Log(config)

	test.Nil(t,err)
	test.NotNil(t,config)
	test.Equal(t,"test",config.Zk)
}

func TestCreateAppConfigWithJsonError(t *testing.T) {
	jsonStr:=`package agollo

import (
	"os"
	"strconv"
	"github.com/cihub/seelog"
	"time"
	"fmt"
	"net/url"
)`
	config,err:=createAppConfigWithJson(jsonStr)
	t.Log(err)

	test.NotNil(t,err)
	test.Nil(t,config)
}

func TestCreateAppConfigWithJsonDefault(t *testing.T) {
	jsonStr:=`{
    "zk": "joe",
    "ip": "localhost:9999"
	}`
	config,err:=createAppConfigWithJson(jsonStr)
	t.Log(err)

	test.Nil(t,err)
	test.NotNil(t,config)
	test.Equal(t,"joe",config.Zk)
}

func TestAppConfig_GetZkIps(t *testing.T) {
	config := AppConfig{Zk: "joe,kkk"}
	ips := config.GetZkIps()

	ipList:=list.New()
	for _,v := range ips{
		ipList.PushBack(v)
	}


	test.HasStringItems(ipList,"joe",t)
}