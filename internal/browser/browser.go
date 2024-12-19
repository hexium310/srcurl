package browser

import "os/exec"

func Open(url string) error {
	return Command(url).Start()
}

func WindowsCommand(url string) *exec.Cmd {
	return exec.Command("cmd.exe", "/c", "start", url)
}

func MacCommand(url string) *exec.Cmd {
	return exec.Command("open", url)
}

func LinuxCommand(url string) *exec.Cmd {
	return exec.Command("xdg-open", url)
}

func WslCommand(url string) *exec.Cmd {
	wslview, err := exec.LookPath("wslview")
	if err != nil {
		return WindowsCommand(url)
	}

	return exec.Command(wslview, url)
}
