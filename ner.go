package main

import (
	"fmt"
	"strings"
)

func main(){
var str = "how do you do"
array := strings.Split(str, " ")
fmt.Println(array)
countMap := make(map[string]int)
for _, item := range array {
countMap[item]++
}
fmt.Println(countMap)
}