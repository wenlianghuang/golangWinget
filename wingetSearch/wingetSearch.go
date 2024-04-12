package wingetsearch

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type App struct {
	Name    string `json:"Name"`
	Id      string `json:"Id"`
	Version string `json:"Version"`
	Match   string `json:"Match,omitempty"`
	Source  string `json:"Source,omitempty"`
}

func WingetJSONSearch() {
	cmd := exec.Command("winget", "search", "--moniker=python", "--source=winget")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	var apps []App
	lines := strings.Split(string(output), "\n")[2:] // Skip header lines

	// Regular expression to match fields separated by one or more spaces
	re := regexp.MustCompile(`^([^ ]+(?: [^ ]+)*) +([^ ]+) +([^ ]+) +([^ ]+(?: [^ ]+)*) +([^ ]+)`)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) == 6 {
			app := App{
				Name:    match[1],
				Id:      match[2],
				Version: match[3],
				Match:   match[4],
				Source:  match[5],
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
	file, err := os.Create("winget_apps.json")
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

	fmt.Println("JSON data written to winget_apps.json successfully")
}
