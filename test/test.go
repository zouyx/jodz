package test

import (
	"testing"
	"reflect"
	"path/filepath"
	"runtime"
	"strings"
	"container/list"
)

func Equal(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("\033[31m%s:%d:\n\n\t   %#v (expected)\n\n\t!= %#v (actual)\033[39m\n\n",
			filepath.Base(file), line, expected, actual)
		t.FailNow()
	}
}

func StartWith(t *testing.T, expected, actual string) {
	Equal(t,true,strings.HasPrefix(actual,expected))
}

func NotEqual(t *testing.T, expected, actual interface{}) {
	if reflect.DeepEqual(expected, actual) {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("\033[31m%s:%d:\n\n\tvalue should not equal %#v\033[39m\n\n",
			filepath.Base(file), line, actual)
		t.FailNow()
	}
}

func Nil(t *testing.T, object interface{}) {
	if !isNil(object) {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("\033[31m%s:%d:\n\n\t   <nil> (expected)\n\n\t!= %#v (actual)\033[39m\n\n",
			filepath.Base(file), line, object)
		t.FailNow()
	}
}

func NotNil(t *testing.T, object interface{}) {
	if isNil(object) {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("\033[31m%s:%d:\n\n\tExpected value not to be <nil>\033[39m\n\n",
			filepath.Base(file), line, object)
		t.FailNow()
	}
}

func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}

func NotEmptyList(list *list.List,t *testing.T)  {
	if list.Len()==0{
		_, file, line, _ := runtime.Caller(1)
		t.Logf("\033[31m%s:%d:\n\n\t   not empty (expected)\n\n",
			filepath.Base(file), line)
		t.FailNow()
	}
}

func HasStringItems(list *list.List,obj interface{},t *testing.T)  {
	NotEmptyList(list,t)

	flag:=false

	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value==obj{
			flag=true
			break
		}
	}

	if !flag{
		_, file, line, _ := runtime.Caller(1)
		t.Logf("\033[31m%s:%d:\n\n\t   contain %#v (expected)\n\n\t not contain %#v (actual)\033[39m\n\n",
			filepath.Base(file), line, obj, obj)
		t.FailNow()
	}
}
