package jodz

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"github.com/zouyx/jodz/config"
	"github.com/zouyx/jodz/utils"
	"github.com/cihub/seelog"
	"fmt"
	"strings"
)

var(
	conn *zk.Conn
)

type Job struct {
	JobName string `json:"jobName"`
}

func init() {
	//fist establishes connection
	connect(config.GetAppConfig())

	//init node
	createParentNode(config.GetAppConfig())
}

//do connect with app config
func connect(appConfig *config.AppConfig)  {
	var err error
	conn,_,err=zk.Connect(appConfig.GetZkIps(), time.Second) //*10)
	if err != nil {
		panic(err)
	}
}

//create job's parent node for create job node
func createParentNode(appConfig *config.AppConfig) {

	if utils.IsEmpty(appConfig.Jobs){
		panic("jobs must config!example: job1(name),jobs2(name)")
	}

	jobs := strings.Split(appConfig.Jobs, comma)

	//init parent nodes
	for _,node:=range parentNodes {
		s, e := conn.Create(node, []byte(""),0, zk.WorldACL(zk.PermAll))

		if utils.IsNotNil(e){
			seelog.Warnf("Connect zk Server Fail,node:%s,Error:%s,",node,e)
			continue
		}

		seelog.Infof("create node:%s success! msg:%s",node,s)
	}

	//init job nodes
	for _,node:=range jobs {
		s, e := conn.Create(getNodeName(jobTemplate,node), []byte(""),0, zk.WorldACL(zk.PermAll))

		if utils.IsNotNil(e){
			seelog.Warnf("Connect zk Server Fail,node:%s,Error:%s,",node,e)
			continue
		}

		seelog.Infof("create node:%s success! msg:%s",node,s)
	}
}

//get node by template and jobName
func getNodeName(template,jobName string) string{
	return fmt.Sprintf(template,jobName)
}

//create job node in zk by job name
func CreateJobNode(jobName string){
	node := getNodeName(ipTemplate, jobName)

	s, e := conn.Create(node, getJobInfo(jobName), zk.FlagEphemeral,zk.WorldACL(zk.PermAll))

	if utils.IsNotNil(e){
		seelog.Warnf("Connect zk Server Fail,node:%s,Error:%s,",node,e)
		return
	}

	seelog.Info("return str:"+s)
}

// get zk connection
func getConn() *zk.Conn{
	return conn
}