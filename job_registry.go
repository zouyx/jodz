package jodz

import (
	"github.com/zouyx/jodz/utils"
	"encoding/json"
	"github.com/cihub/seelog"
)

var(
	jobMap=make(map[string]JobScheduler,0)

	jobInfoMap=make(map[string][]byte,0)
)

//registry job by name
func RegistryJob(jobName string,scheduler JobScheduler)  {
	jobMap[jobName]=scheduler
}



func getJobInfo(jobName string) []byte {
	jobJson := jobInfoMap[jobName]

	if utils.IsNotNil(jobJson){
		return jobJson
	}

	job := &Job{
		JobName: jobName,
	}

	bytes, e := json.Marshal(job)
	if utils.IsNotNil(e){
		seelog.Error("json format fail!error:",e)
		return []byte{}
	}

	jobInfoMap[jobName]= bytes

	return bytes
}

//run job by name
//return
//true: run success
//false: run fail
func RunJob(jobName string)bool{
	scheduler:=jobMap[jobName]
	if utils.IsNil(scheduler){
		return false
	}

	scheduler.RunTask()

	return true
}

