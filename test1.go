package main

import (
	"fmt"
	"reflect"
	"study/aaaa"
)

type idao interface {
	xxx(string2 string) string
}


func (p person) xxx(s string) string {
	return "xxx"
}
type person struct{
	id int
	name string
}
func main() {
	m := aaaa.Mmm{Id: 123}

	print(m.Id)
	a := new(aaaa.Mmm)
	fmt.Println(reflect.TypeOf(a))
}
