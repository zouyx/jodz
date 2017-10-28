package jodz

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"github.com/zouyx/jodz/config"
	"github.com/zouyx/jodz/utils"
)
const(
	prefix="/jodz"

	jobScheduler=prefix+"/jobScheduler"
)

var(
	conn *zk.Conn
)

func init() {
	//fist establishes connection
	connect(config.GetAppConfig())
}

func connect(appConfig *config.AppConfig)  {
	var err error
	conn,_,err=zk.Connect(appConfig.GetZkIps(), time.Second) //*10)
	if err != nil {
		panic(err)
	}
}

func CreateNode(jobName string){
	s, e := conn.CreateProtectedEphemeralSequential(jobScheduler, []byte(utils.GetInternal()), zk.WorldACL(zk.PermAll))

	if utils.IsNotNil(e){

	}
}