package execution

import (
	"os"
	"os/exec"
	"runtime"
)

func ExecuteExternalCommand(args []string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// On Windows, use `cmd.exe` to execute commands
		cmd = exec.Command("cmd", append([]string{"/C"}, args...)...)
	} else {
		// On Unix-like systems, use `sh -c` for command execution
		cmd = exec.Command("sh", "-c", args[0])
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
