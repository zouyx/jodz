package registry

import "github.com/jodz/utils"

var(
	jobMap map[string]JobScheduler
)

//registry job by name
func RegistryJob(jobName string,scheduler JobScheduler)  {
	jobMap[jobName]=scheduler
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

