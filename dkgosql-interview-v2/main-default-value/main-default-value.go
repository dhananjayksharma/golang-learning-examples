package main

import "fmt"

func main() {
	var name *string
	fmt.Println("name as string pointer is <nil>:", name)
	var namestr string
	fmt.Println("name as string is '':", namestr)

	var year *int
	fmt.Println("year as int pointer is <nil>:", year)

	var yearasint int
	fmt.Println("year as int is 0:", yearasint)

	var coursepoint map[string]string
	// coursepoint["bca"] = "bachelor"
	fmt.Println("coursepoint as map is nil:", coursepoint == nil)

	var course = make(map[string]string)
	course["bca"] = "bachelor"
	fmt.Println("coursepoint as map is with data:", course)

	var courses []string
	fmt.Println("courses as slices is nil:", courses == nil)

	var coursesdata = []string{"str", "int"}
	fmt.Println("coursesdata as slices is with data:", coursesdata)

}
