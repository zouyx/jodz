package jodz

import (
	"github.com/cihub/seelog"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/zouyx/jodz/config"
)

func init() {
	appConfig := config.GetAppConfig()

	//init node for config
	for _,jobNode := range appConfig.GetJobNodes() {
		//create job node
		CreateJobNode(jobNode)

		//watch job node
		go watchJob(jobNode,nil,nil)
	}
}

func watchJob(jobName string,
	succCallBack func(zk.Event, []byte),
	failCallBack func(error)) {

	conn := getConn()
	data, _, events, e := conn.GetW(getNodeName(jobTemplate, jobName))

	if e != nil {
		seelog.Errorf("watch job:%s fail!error:%s", jobName, e)
		if failCallBack != nil {
			failCallBack(e)
			return
		}
	}
	ev := <-events
	//
	if succCallBack==nil{
		seelog.Warnf("watch job:%s trigger success callback fail! because no function.")
	}
	succCallBack(ev, data)
}

//dead loop for watchJob Node
func loopWatchJob(jobName string,
	succCallBack func(zk.Event, []byte),
	failCallBack func(error)) {
	for  {
		watchJob(jobName,succCallBack,failCallBack)
	}
}
