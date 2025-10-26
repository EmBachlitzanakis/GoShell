package execution

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ExecuteExternalCommand(args []string) error {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		// On Windows, use `cmd.exe` to execute commands
		cmdArgs := append([]string{"/C"}, args...)
		cmd = exec.Command("cmd", cmdArgs...)
	} else {
		// On Unix-like systems, use `sh -c` for command execution
		cmd = exec.Command(args[0], args[1:]...)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute command '%s': %w", strings.Join(args, " "), err)
	}

	return nil
}
