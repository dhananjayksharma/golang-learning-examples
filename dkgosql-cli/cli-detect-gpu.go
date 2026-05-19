package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("system_profiler", "SPDisplaysDataType")

	case "windows":
		cmd = exec.Command("powershell", "-Command",
			"Get-CimInstance Win32_VideoController | Select-Object Name,AdapterRAM,DriverVersion")

	case "linux":
		cmd = exec.Command("sh", "-c", "lspci | grep -i 'vga\\|3d\\|display'")

	default:
		fmt.Println("Unsupported OS:", runtime.GOOS)
		os.Exit(1)
	}

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
