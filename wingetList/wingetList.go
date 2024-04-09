package wingetlist

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type App struct {
	Name    string `json:"Name"`
	Id      string `json:"Id"`
	Version string `json:"Version"`
	Match   string `json:"Match"`
	Source  string `json:"Source"`
}

func WingetJSONList() {
	cmd := exec.Command("winget", "search", "--query=python3.10")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	var apps []App
	lines := strings.Split(string(output), "\n")[2:] // Skip header lines

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) >= 7 {
			//have space in name
			name := strings.Join(fields[:2], " ")
			//have space in match
			match := strings.Join(fields[4:6], " ")
			app := App{
				Name:    name,
				Id:      fields[2],
				Version: fields[3],
				Match:   match,
				Source:  fields[6],
			}
			apps = append(apps, app)
		}
	}

	jsonData, err := json.MarshalIndent(apps, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	//fmt.Println(string(jsonData))
	jsonStr := strings.TrimPrefix(string(jsonData), "[")
	jsonStr = strings.TrimSuffix(jsonStr, "]")

	file, err2 := os.Create("winget_apps.json")
	if err2 != nil {
		fmt.Println("Error to create file: ", err2)
		return
	}
	defer file.Close()
	_, err2 = file.WriteString(jsonStr)
	if err2 != nil {
		fmt.Println("Error writing to file: ", err2)
		return
	}
	fmt.Println("JSON data written to winget_app.json successfully")
}
