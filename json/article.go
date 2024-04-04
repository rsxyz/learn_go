package main

import (
	"encoding/json"
	"fmt"
)

type Article struct {
	Id      string
	Title   string
	Content string
	Summary string
}

func main() {
	var post Article
	article := `{"id": "BM-1347", "title": "The underage storm", "Content": "The creatives' careers can easily get uncreative but yet creative...", "Summary": "Seeking freedom"}`

	json.Unmarshal([]byte(article), &post)

	fmt.Printf("Recently posted article: %s", post)

	newData, err := json.Marshal(post)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newData))
	}
}
