package wingetlist

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type AppTag struct {
	Name    string `json:"Name"`
	Id      string `json:"Id"`
	Version string `json:"Version"`
	Source  string `json:"Source"`
}

func WingetListtag() {
	cmd := exec.Command("winget", "search", "Golang.Go")
	//cmd := exec.Command("winget", "search", "--id=Python.Python")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	var apps []AppTag
	lines := strings.Split(string(output), "\n")[2:] // Skip header lines

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) >= 6 {
			//have space in name
			name := strings.Join(fields[:3], " ")
			app := AppTag{
				Name:    name,
				Id:      fields[3],
				Version: fields[4],
				Source:  fields[5],
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
