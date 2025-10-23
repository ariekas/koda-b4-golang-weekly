package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windos":
		cmd = exec.Command("cmd", "/c", "cls")
	default: 
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
