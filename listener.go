package jodz

import (
	"github.com/cihub/seelog"
)

func watchJob(jobName string)  {
	conn := getConn()
	bytes, stat, events, e := conn.GetW(getNodeName(jobTemplate, jobName))
	seelog.Info("e:",e)
	seelog.Info("stat:",stat)
	seelog.Info("bytes:",string(bytes))

	ev:=<-events
	seelog.Info("events:",ev)
}