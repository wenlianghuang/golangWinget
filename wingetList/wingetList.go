package wingetlist

import (
	"fmt"
	"os/exec"
)

func WingetList() {
	cmd := exec.Command("winget", "list", "--id=python")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("No application in your computuer, please search in the web")
		cmdsearch := exec.Command("winget", "search", "--query=python3.10")
		output2, err2 := cmdsearch.Output()
		if err2 != nil {
			fmt.Println("Command Error: ", err2)
		}
		fmt.Println(string(output2))
		return
	}
	fmt.Println(string(output))
}
