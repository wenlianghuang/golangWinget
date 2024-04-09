package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("winget", "list", "--id=python")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Println(string(output))

	cmddownload := exec.Command("winget", "uninstall", "--id=Python.Python.3.11")
	_, err2 := cmddownload.Output()
	if err2 != nil {
		fmt.Println("Error2 executing command:", err2)
		return
	}
	fmt.Println("Uninstall Python3.11  successfully")
}
