package main

import (
	"encoding/json"
	"fmt"
)

type CommentList struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID     int        `json:"id"`
	Body   string     `json:"body"`
	PostID int        `json:"postId"`
	User   UserStruct `json:"user"`
}
type UserStruct struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func ShowMostActiveUser(comment CommentList) {
	var mostActiveUser = make(map[int]int)

	for _, v := range comment.Comments {
		userID := v.User.ID
		if _, ok := mostActiveUser[userID]; !ok {
			mostActiveUser[userID] = 1
		} else {
			mostActiveUser[userID] += 1
		}
	}
	mostActiveUserKey := make([]int, 0, len(mostActiveUser))

	for _, v := range mostActiveUser {
		mostActiveUserKey = append(mostActiveUserKey, v)
	}

	for i := 0; i < len(mostActiveUserKey); i++ {
		min_idx := i
		for j := i + 1; j < len(mostActiveUserKey); j++ {
			if mostActiveUserKey[min_idx] < mostActiveUserKey[j] {
				min_idx = j
			}
		}
		mostActiveUserKey[i], mostActiveUserKey[min_idx] = mostActiveUserKey[min_idx], mostActiveUserKey[i]
	}
	fmt.Println("Before data mostActiveUser:", mostActiveUser)
	fmt.Println("Before keys mostActiveUserKey:", mostActiveUserKey)
	for _, cntNum := range mostActiveUserKey {
		id := getDataByCountValue(cntNum, mostActiveUser)
		fmt.Printf("name:%d, cnt:%d\n", id, cntNum)
	}

}

func getDataByCountValue(cntNum int, mostActiveUser map[int]int) int {
	for key, v := range mostActiveUser {
		if v == cntNum {
			delete(mostActiveUser, key)
			return key
		}
	}
	return 0
}
func main() {
	fmt.Println("start main")
	jsonData := `{"comments":[{"id":1,"body":"This is some awesome thinking!","postId":100,"user":{"id":63,"username":"eburras1q"}},{"id":2,"body":"What terrific math skills you’re showing!","postId":27,"user":{"id":71,"username":"omarsland1y"}},{"id":3,"body":"You are an amazing writer!","postId":61,"user":{"id":29,"username":"jissetts"}},{"id":4,"body":"Wow! You have improved so much!","postId":8,"user":{"id":19,"username":"bleveragei"}},{"id":5,"body":"Nice idea!","postId":62,"user":{"id":70,"username":"cmasurel1x"}},{"id":6,"body":"You are showing excellent understanding!","postId":19,"user":{"id":97,"username":"cdavydochkin2o"}},{"id":7,"body":"This is clear, concise, and complete!","postId":47,"user":{"id":22,"username":"froachel"}},{"id":8,"body":"What a powerful argument!","postId":47,"user":{"id":82,"username":"kogilvy29"}},{"id":9,"body":"I knew you could do it!","postId":64,"user":{"id":31,"username":"smargiottau"}},{"id":10,"body":"Wonderful ideas!","postId":4,"user":{"id":35,"username":"mbrooksbanky"}},{"id":11,"body":"It was a pleasure to grade this!","postId":2,"user":{"id":68,"username":"rstrettle1v"}},{"id":12,"body":"Keep up the incredible work!","postId":50,"user":{"id":77,"username":"rkingswood24"}},{"id":13,"body":"My goodness, how impressive!","postId":37,"user":{"id":28,"username":"xisherwoodr"}},{"id":14,"body":"You’re showing inventive ideas!","postId":30,"user":{"id":57,"username":"bpickering1k"}},{"id":15,"body":"You’ve shown so much growth!","postId":44,"user":{"id":76,"username":"cgaber23"}},{"id":16,"body":"Interesting thoughts!","postId":71,"user":{"id":100,"username":"pcumbes2r"}},{"id":17,"body":"I love your neat work!","postId":68,"user":{"id":37,"username":"nwytchard10"}},{"id":18,"body":"Doesn’t it feel good to do such great work?","postId":41,"user":{"id":31,"username":"smargiottau"}},{"id":19,"body":"First-rate work!","postId":75,"user":{"id":60,"username":"dlambarth1n"}},{"id":20,"body":"This is fascinating information!","postId":48,"user":{"id":17,"username":"vcholdcroftg"}},{"id":21,"body":"You inspire me!","postId":29,"user":{"id":5,"username":"kmeus4"}},{"id":22,"body":"This is right on target!","postId":18,"user":{"id":31,"username":"smargiottau"}},{"id":23,"body":"What an astounding observation!","postId":73,"user":{"id":14,"username":"mturleyd"}},{"id":24,"body":"This is very well thought out!","postId":32,"user":{"id":16,"username":"dpierrof"}},{"id":25,"body":"I can tell you’ve been practicing!","postId":44,"user":{"id":78,"username":"dbuist25"}},{"id":26,"body":"You’ve come a long way!","postId":70,"user":{"id":82,"username":"kogilvy29"}},{"id":27,"body":"I can tell you’ve been paying attention!","postId":60,"user":{"id":74,"username":"ahinckes21"}},{"id":28,"body":"Reading this made my day!","postId":85,"user":{"id":85,"username":"kpondjones2c"}},{"id":29,"body":"This is very perceptive!","postId":13,"user":{"id":30,"username":"kdulyt"}},{"id":30,"body":"What an accomplishment!","postId":23,"user":{"id":36,"username":"dalmondz"}}]}`
	var comment CommentList
	err := json.Unmarshal([]byte(jsonData), &comment)

	if err == nil {
		fmt.Println("Error found:", err)
	}

	ShowMostActiveUser(comment)
}
