package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type VersionInfo struct {
	ServerVersion struct {
		Major string `json:"major"`
		Minor string `json:"minor"`
	} `json:"serverVersion"`
}

func main() {
	cmd := exec.Command("kubectl", "version", "-o", "json")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing kubectl version command:", err)
		return
	}

	var version VersionInfo
	err = json.Unmarshal(out, &version)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	serverMajor := version.ServerVersion.Major
	serverMinor := version.ServerVersion.Minor

	fmt.Println("Server Version Major:", serverMajor)
	fmt.Println("Server Version Minor:", serverMinor)
}
