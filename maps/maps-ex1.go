package main

import "fmt"

func main(){
	states :=  map[string]string{
		"RJ":"Rajasthaan",
		"BH":"Bihar",
		"JH":"Jharkhand",
		"HP":"Himachal Pradesh",
	}

	for i,v := range states{
		fmt.Printf("key:%v, val:%v\n", i, v)
	}
}