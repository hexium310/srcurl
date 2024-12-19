//go:build windows

package browser

import (
	"os/exec"
	"syscall"
)

func Command(url string) *exec.Cmd {
	command := WindowsCommand(url)
	command.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	return command
}
