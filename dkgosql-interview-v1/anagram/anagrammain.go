package main

import "fmt"

func checkAnagram(astr, bstr string) {
	var outmap = make(map[int]string, len(astr))
	for i := 0; i < len(astr); i++ {
		outmap[i] = string(astr[i])
	}

	allcount := len(astr)
	searchcnt := 0
	for i := 0; i < len(bstr); i++ {
		out := string(bstr[i])
		outstatus := checkinmap(outmap, out)
		if outstatus {
			searchcnt++
		}
	}
	if allcount == searchcnt {
		fmt.Println("its anagram")
	} else {
		fmt.Println("Not anagram")
	}
}
func checkinmap(outmap map[int]string, out string) bool {
	for _, v := range outmap {
		if v == out {
			return true
		}
	}
	return false
}

func main() {

	astr := "earth"
	bstr := "heart"

	checkAnagram(astr, bstr)
}

// 112233
