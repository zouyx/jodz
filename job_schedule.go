package jodz

import "github.com/samuel/go-zookeeper/zk"

type JobScheduler interface {
	RunTask()

	SuccessCallBack(zk.Event, []byte)

	FailCallBack(error)

}
