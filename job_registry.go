package jodz

import (
	"github.com/zouyx/jodz/utils"
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/samuel/go-zookeeper/zk"
)

var(
	jobMap=make(map[string]JobScheduler,0)

	jobInfoMap=make(map[string][]byte,0)

	//cache job name
	jobNames=make([]string,0)
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

func GetJobNodes() map[string]JobScheduler{
	return jobMap
}

//get all job names
func GetJobNames() []string{
	if jobNames!=nil&&len(jobNames)>0{
		return jobNames
	}

	for jobName, _ := range jobMap {
		jobNames=append(jobNames,jobName)
	}

	return jobNames
}

func RegistryAllJobs()  {
	jobs:=GetJobNames()

	if len(jobs)==0{
		panic("jobs must config!example: job1(name),jobs2(name)")
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