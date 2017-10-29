package jodz

import "github.com/zouyx/jodz/utils"

const(
	//parent nodes
	prefix="/jodz"

	jobScheduler=prefix+"/jobScheduler"

	jobTemplate=jobScheduler+"/%s"

	//split symbol
	comma=","
)

var(
	parentNodes=[]string{prefix,jobScheduler}

	ipTemplate=jobTemplate+"/"+utils.GetInternal()
)