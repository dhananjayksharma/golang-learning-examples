package main

import "fmt"

func main(){
	i := 6
	var k int = 0
	for j:=0; j<i; j++{
	 k += j	
		fmt.Println(j, k)
	}
}