package main

import (
	"fmt"
	"strings"
)

type MapSum struct {
	m map[string]int
}


/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		m: map[string]int{},
	}
}


func (this *MapSum) Insert(key string, val int)  {
	this.m[key]=val
}


func (this *MapSum) Sum(prefix string) int {
	s:=0
	for k,v:=range this.m{
		if strings.HasPrefix(k,prefix) {
			s+=v
		}
	}
	return s
}

func main()  {
	key:="apple"
	val:=3
	obj := Constructor();
	obj.Insert(key,val);
	obj.Insert("app",2);
	param_2 := obj.Sum("app");
	fmt.Println(param_2)
}

