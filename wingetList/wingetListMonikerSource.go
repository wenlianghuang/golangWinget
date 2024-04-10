package wingetlist

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type AppMonkirTag struct {
	Name    string `json:"Name"`
	Id      string `json:"Id"`
	Version string `json:"Version"`
	//Match   string `json:"Match,omitempty"`
}

func WingetMonikerSource() {
	cmd := exec.Command("winget", "search", "--moniker=vs2019", "--source=winget")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	var apps []AppMonkirTag
	lines := strings.Split(string(output), "\n")[2:] // Skip header lines

	// Regular expression to match fields separated by one or more spaces
	re := regexp.MustCompile(`^([^ ]+(?: [^ ]+)*) +([^ ]+) +([^ ]+) +([^ ]+(?: [^ ]+)*) +([^ ]+)`)
	//re := regexp.MustCompile(`^([^ ]+(?: [^ ]+)*) +([^ ]+) +([^ ]+) +(Moniker: .+?) +([^ ]+)`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) == 6 {
			app := AppMonkirTag{
				Name:    match[1],
				Id:      match[2],
				Version: match[3],
				//Match:   match[4],
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
