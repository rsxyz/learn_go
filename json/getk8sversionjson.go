package main

import (
	"encoding/json"
	"fmt"
)

type VersionInfo struct {
	ServerVersion struct {
		Major string `json:"major"`
		Minor string `json:"minor"`
	} `json:"serverVersion"`
}

func main() {
	jsonStr := `{
		"clientVersion": {
			"major": "1",
			"minor": "25",
			"gitVersion": "v1.25.2",
			"gitCommit": "5835544ca568b757a8ecae5c153f317e5736700e",
			"gitTreeState": "clean",
			"buildDate": "2022-09-21T14:33:49Z",
			"goVersion": "go1.19.1",
			"compiler": "gc",
			"platform": "darwin/amd64"
		},
		"kustomizeVersion": "v4.5.7",
		"serverVersion": {
			"major": "1",
			"minor": "25",
			"gitVersion": "v1.25.3",
			"gitCommit": "434bfd82814af038ad94d62ebe59b133fcb50506",
			"gitTreeState": "clean",
			"buildDate": "2022-10-25T19:35:11Z",
			"goVersion": "go1.19.2",
			"compiler": "gc",
			"platform": "linux/amd64"
		}
	}`

	var v VersionInfo
	if err := json.Unmarshal([]byte(jsonStr), &v); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	serverMajor := v.ServerVersion.Major
	serverMinor := v.ServerVersion.Minor

	fmt.Printf("Server Version Major: %s\n", serverMajor)
	fmt.Printf("Server Version Minor: %s\n", serverMinor)
}
