package jodz

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"github.com/zouyx/jodz/config"
	"github.com/zouyx/jodz/utils"
	"github.com/cihub/seelog"
)
const(
	prefix="/jodz"

	jobScheduler=prefix+"/jobScheduler"


)

var(
	conn *zk.Conn

	ip=jobScheduler+"/"+utils.GetInternal()
)

func init() {
	//fist establishes connection
	connect(config.GetAppConfig())

	//init node
	createParentNode()
}

func connect(appConfig *config.AppConfig)  {
	var err error
	conn,_,err=zk.Connect(appConfig.GetZkIps(), time.Second) //*10)
	if err != nil {
		panic(err)
	}
}

func createParentNode() {

	s, e := conn.Create(prefix, []byte(""),0, zk.WorldACL(zk.PermAll))

	if utils.IsNotNil(e){
		seelog.Error("Connect zk Server Fail,Error:",e)
		return
	}

	seelog.Info("return str:"+s)

	s, e = conn.Create(jobScheduler, []byte(""),0, zk.WorldACL(zk.PermAll))

	if utils.IsNotNil(e){
		seelog.Error("Connect zk Server Fail,Error:",e)
		return
	}

	seelog.Info("return str:"+s)
}

func CreateNode(jobName string){
	s, e := conn.Create(ip, []byte(""), zk.FlagEphemeral,zk.WorldACL(zk.PermAll))

	if utils.IsNotNil(e){
		seelog.Error("Connect zk Server Fail,Error:",e)
		return
	}

	seelog.Info("return str:"+s)
}