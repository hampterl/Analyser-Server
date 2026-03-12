package internal

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func EnsureDirs(base string) error {
	dirs := []string{
		base,
		base + "\\bin",
		base + "\\logs",
		base + "\\data",
		base + "\\config",
	}

	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}
	return nil
}

func CurrentExe() (string, error) {
	return os.Executable()
}

func CreateStartupLink(name, exePath string) error {
	startup := filepath.Join(
		os.Getenv("APPDATA"),
		"Microsoft", "Windows", "Start Menu", "Programs", "Startup",
	)

	lnkPath := filepath.Join(startup, name+".lnk")
	workDir := filepath.Dir(exePath)

	ps := fmt.Sprintf(`
$WshShell = New-Object -ComObject WScript.Shell
$Shortcut = $WshShell.CreateShortcut("%s")
$Shortcut.TargetPath = "%s"
$Shortcut.WorkingDirectory = "%s"
$Shortcut.Save()
`, lnkPath, exePath, workDir)

	cmd := exec.Command("powershell", "-NoProfile", "-Command", ps)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	return cmd.Run()
}

func Relaunch(target string) error {
	cmd := exec.Command(target)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

func CopySelf(target string) error {
	src, err := os.Executable()
	if err != nil {
		return err
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
