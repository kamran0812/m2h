package internal

import (
	"fmt"
	"os/exec"
	"runtime"
)

func Preview(fName string) error {
	command := ""
	param := []string{}

	switch runtime.GOOS {
	case "linux":
		command = "xdg-open"
	case "windows":
		command = "cmd.exe"
		param = []string{"/C", "start"}
	case "darwin":
		command = "open"
	default:
		return fmt.Errorf("OS not supported")

	}
	param = append(param, fName)

	cmdPath, err := exec.LookPath(command)
	if err != nil {
		return err
	}

	return exec.Command(cmdPath, param...).Run()

}
