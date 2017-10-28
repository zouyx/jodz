package jodz

import (
	"testing"
	"time"
)

func TestCreateNode(t *testing.T) {
	CreateNode("zou")

	time.Sleep(50*time.Second)
}
