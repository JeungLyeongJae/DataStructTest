package main

import (
	"DataStructTest/model/data_struct/list"
	"fmt"
)

func main() {
	l := &list.LinkedList{}
	l.Insert("1")
	l.Insert("2")
	l.Insert("3")
	l.Insert("4")
	l.Insert("5")

	fmt.Print(l.IndexOf("5"))
}
