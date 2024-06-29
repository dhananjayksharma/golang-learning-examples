package main

import "fmt"

type X interface {
	~int | float64 | float32 | ~string
}

func show[T X](list []T) {
	for c, v := range list {
		fmt.Println(c, v)
	}
}
func main() {
	x := []int{1, 2, 5}
	show(x)
	f := []float64{1.3, 2.6, 5.9}
	show(f)
	s := []string{"hi", "golang"}
	show(s)
	sli := []map[string][]string{{
		"days": {"sun", "mon"},
		"week": {"weekdays"},
	}}
	fmt.Println(sli)
}
