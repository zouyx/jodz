package registry

var(
	jobMap map[string]JobScheduler
)

//registry job by name
func RegistryJob(jobName string,scheduler JobScheduler)  {
	jobMap[jobName]=scheduler
}


