//go:build !windows

package browser

import (
	"os/exec"
	"runtime"
	"strings"
)

func Command(url string) *exec.Cmd {
	switch runtime.GOOS {
	case "darwin":
		return MacCommand(url)
	default:
		uname, err := exec.Command("uname", "-r").Output()
		if err != nil {
			panic(err)
		}

		if strings.Contains(string(uname), "microsoft") {
			return WslCommand(url)
		} else {
			return LinuxCommand(url)
		}
	}
}
