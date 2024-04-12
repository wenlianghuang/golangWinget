package wingetsearch

import (
	"fmt"
	"os/exec"
	"strings"
)

type AppNormal struct {
	Name    string
	Id      string
	Version string
}

func WingetNormalSearch() {
	cmd := exec.Command("winget", "search", "--id=python3.10")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	var apps []AppNormal
	lines := strings.Split(string(output), "\n")[2:] // Skip header lines

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 3 {
			app := AppNormal{
				Name:    strings.Join(fields[:len(fields)-2], " "),
				Id:      fields[len(fields)-2],
				Version: fields[len(fields)-1],
			}
			apps = append(apps, app)
		}
	}

	// Find max width for each column
	maxNameWidth := 0
	maxIdWidth := 0
	maxVersionWidth := 0
	for _, app := range apps {
		if len(app.Name) > maxNameWidth {
			maxNameWidth = len(app.Name)
		}
		if len(app.Id) > maxIdWidth {
			maxIdWidth = len(app.Id)
		}
		if len(app.Version) > maxVersionWidth {
			maxVersionWidth = len(app.Version)
		}
	}

	// Print headers
	fmt.Printf("%-*s  %-*s  %-*s\n", maxNameWidth, "Name", maxIdWidth, "Id", maxVersionWidth, "Version")

	// Print data
	for _, app := range apps {
		fmt.Printf("%-*s  %-*s  %-*s\n", maxNameWidth, app.Name, maxIdWidth, app.Id, maxVersionWidth, app.Version)
	}
}
