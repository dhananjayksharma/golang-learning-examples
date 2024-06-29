package main

import (
	"encoding/json"
	"fmt"
)

type ShortenURL struct {
	ShortURL string
	Count    int
}

func main() {
	var mapData = map[string]ShortenURL{}
	var outData = map[string]ShortenURL{}
	mapData["url1"] = ShortenURL{ShortURL: "url1", Count: 1}
	mapData["url1"] = ShortenURL{ShortURL: "url1", Count: 2}
	mapData["url2"] = ShortenURL{ShortURL: "url2", Count: 1}

	byteData, err := json.Marshal(mapData)
	if err != nil {
		fmt.Printf("Error while marshal data:%v", err)
	}
	fmt.Println("JsonData:", string(byteData))

	err = json.Unmarshal(byteData, &outData)
	if err != nil {
		fmt.Printf("Error while marshal data:%v", err)
	}
	fmt.Println("outData url1:", outData["url1"].ShortURL, outData["url1"].Count)
	fmt.Println("outData url2:", outData["url2"].ShortURL, outData["url2"].Count)
}
