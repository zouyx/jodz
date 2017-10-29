package jodz

import (
	"testing"
	"time"
	"github.com/zouyx/jodz/test"
)

func TestCreateNode(t *testing.T) {
	joeName := "zou"
	CreateJobNode(joeName)

	time.Sleep(2*time.Second)

	bytes, _, e := conn.Get(getNodeName(ipTemplate, joeName))

	test.Equal(t,`{"jobName":"zou"}`,string(bytes))
	test.Nil(t,e)
}
