package jodz

import (
	"testing"
	"github.com/cihub/seelog"
	"time"
	"github.com/samuel/go-zookeeper/zk"
)

type JoeJobScheduler struct {

}

func (this JoeJobScheduler)SuccessCallBack(e zk.Event, b []byte) {
	seelog.Info("data:",string(b))
	seelog.Info("event:",e)
}

func (this JoeJobScheduler)FailCallBack(e error) {
	seelog.Info("error:",e)
}

func (this JoeJobScheduler)RunTask() {
}


func TestWatchJob(t *testing.T) {
	scheduler := &JoeJobScheduler{}

	RegistryJob("joe",scheduler)

	RegistryAllJobs()


	jobName := getNodeName(jobTemplate, "joe")

	go setData(jobName)

	watchJob("joe",
		func(event zk.Event, bytes []byte) {
		seelog.Infof("success,data:%s",string(bytes))
	}, func(e error) {
			seelog.Errorf("fail,data:%s",e)
	})
}

func setData(jobName string) {
	time.Sleep(3*time.Second)
	_, stat, _ := conn.Get(jobName)

	stat, e := conn.Set(jobName,[]byte("gogogo"),stat.Version)
	seelog.Info("stat:",stat)
	seelog.Info("e:",e)
}
