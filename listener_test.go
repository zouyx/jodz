package jodz

import (
	"testing"
	"github.com/cihub/seelog"
	"time"
)

func TestWatchJob(t *testing.T) {
	jobName := getNodeName(jobTemplate, "joe")

	go setData(jobName)

	//watchJob("joe")
}

func setData(jobName string) {
	time.Sleep(3*time.Second)
	stat, e := conn.Set(jobName,[]byte("gogogo"),0)
	seelog.Info("stat:",stat)
	seelog.Info("e:",e)
}
