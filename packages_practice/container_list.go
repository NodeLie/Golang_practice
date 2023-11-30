package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List
	for i := 0; i < 3; i++ {
		x.PushBack(i)
	}

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
