package main

import "fmt"

 

func main(){
	 var listData []int
	 add(&listData)
	 fmt.Println(listData)
}

func add(listData *[]int){
	*listData = append(*listData, 10)
	// listData = append(listData, 13)
}
