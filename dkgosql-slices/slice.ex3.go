package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	max, _ := strconv.Atoi(os.Args[1])
	// const max = 5
	var uniques []int
	// var uniques = make([]int, max)
loop:
	for len(uniques) < max {
		n := rand.Intn(max) + 1
		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}
		uniques = append(uniques, n)
	}
	sort.Ints(uniques)
	fmt.Println("uniques:", uniques)
	fmt.Println("len, cap:", len(uniques), cap(uniques))

	nums := [5]int{5, 8, 4, 7, 5}
	sort.Ints(nums[:])
	fmt.Println("nums:", nums)

	nums2 := []int{9, 7, 5}
	append(nums2, []int{2, 4, 6}...)

	fmt.Println(nums2[3])
}
