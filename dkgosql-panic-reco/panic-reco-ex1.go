package main

import "fmt"

func recoverNow(dividend, devisor int) {
	if r := recover(); r != nil {
		fmt.Printf("RECOVERED FROM dividend:%d / devisor:%d , error:%v\n", dividend, devisor, r)
	}
}

func divide(dividend, devisor int) float64 {
	defer recoverNow(dividend, devisor)

	return float64(dividend / devisor)
}

// nil pointer, resource, too many resource as file opened

func main() {
	fmt.Println("Start execution")
	datalist := map[int]int{100: 0, 12: 4, 34: 3, 23: 0, 45: 9, 90: 0}

	for i, v := range datalist {
		out := divide(i, v)
		fmt.Printf("dividend:%d / devisor:%d : remain:%f\n", i, v, out)
	}
	fmt.Println("main exit")
}
