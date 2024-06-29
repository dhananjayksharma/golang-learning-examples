package main

import "fmt"

func main(){
	fmt.Println("---Table---")
	number := 11

	for i:=1; i<= 5; i++{
		fmt.Println(number, "x",i," = ", number*i)
	}

}