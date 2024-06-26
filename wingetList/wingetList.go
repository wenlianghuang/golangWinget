package wingetlist

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type AppList struct {
	Name    string `json:"Name"`
	Id      string `json:"Id"`
	Version string `json:"Version"`
}

func WingetJSONList() {
	cmd := exec.Command("winget", "list", "Microsoft Visual")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	var apps []AppList
	lines := strings.Split(string(output), "\n")[2:] // Skip header lines

	// Regular expression to match fields separated by one or more spaces
	re := regexp.MustCompile(`^([^ ]+(?: [^ ]+)*) +([^ ]+) +([^ ]+)`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) == 4 {
			app := AppList{
				Name:    match[1],
				Id:      match[2],
				Version: match[3],
			}
			apps = append(apps, app)
		}
	}

	jsonData, err := json.MarshalIndent(apps, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write JSON data to file
	file, err := os.Create("winget_lists.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON data written to winget_lists.json successfully")
}
