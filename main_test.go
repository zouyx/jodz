package jodz

import (
	"testing"
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"fmt"
)

func Test_Main(t *testing.T) {
	c, _, err := zk.Connect([]string{"127.0.0.1:2180"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	children, stat, ch, err := c.ChildrenW("/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", children, stat)
	e := <-ch
	fmt.Printf("%+v\n", e)
}
