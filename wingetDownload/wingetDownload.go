package wingetdownload

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type PythonInfo struct {
	Name    string `json:"Name"`
	Id      string `json:"Id"`
	Version string `json:"Version"`
}

func WingetDownloadfromJSON() {
	// Open the JSON file
	file, err := os.Open("./winget_apps.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode JSON data from the file
	var pythonInfoList []PythonInfo
	err = json.NewDecoder(file).Decode(&pythonInfoList)
	if err != nil {
		log.Fatal(err)
	}

	// Search for Python Launcher
	var infoName string
	found := false
	for _, info := range pythonInfoList {
		if strings.Contains(info.Name, "Visual Studio Professional 2019") {
			fmt.Printf("Name: %s\nId: %s\nVersion: %s\n", info.Name, info.Id, info.Version)
			found = true
			infoName = info.Id
			break
		}
	}

	if !found {
		fmt.Println("Python Launcher not found in the JSON data.")
	}
	downloadid := "--id=" + "\"" + infoName + "\""
	cmd := exec.Command("winget", "download", downloadid)
	_, err2 := cmd.Output()
	//_, err := cmd.Output()
	if err2 != nil {
		fmt.Println("Error executing command:", err)
		return
	}

}
