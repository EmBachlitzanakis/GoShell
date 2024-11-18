package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Error constants
var (
	ErrNoPath = errors.New("path required")
)

func main() {
	startShell()
}

// startShell initializes the shell loop for user input
func startShell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		// Process and execute the input
		if err := handleInput(input); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	}
}

// handleInput processes user input and routes it to the appropriate command
func handleInput(input string) error {
	input = strings.TrimSpace(input)

	// Ignore empty input
	if input == "" {
		return nil
	}

	// Split input into command and arguments
	args := strings.Split(input, " ")
	command := args[0]

	// Check for built-in commands
	if isBuiltInCommand(command) {
		return executeBuiltInCommand(command, args)
	}

	// If not a built-in command, run as external command
	return executeExternalCommand(input, args)
}

// isBuiltInCommand checks if a command is built-in
func isBuiltInCommand(command string) bool {
	switch command {
	case "cd", "exit", "help", "ls", "dir":
		return true
	}
	return false
}

// executeBuiltInCommand executes supported built-in commands
func executeBuiltInCommand(command string, args []string) error {
	switch command {
	case "cd":
		return changeDirectory(args)
	case "exit":
		exitShell()
	case "help":
		printHelp()
		return nil
	case "ls", "dir":
		return handleCrossPlatformDir(args)
	}
	return fmt.Errorf("unknown built-in command: %s", command)
}

// changeDirectory handles the 'cd' command
func changeDirectory(args []string) error {
	if len(args) < 2 {
		return ErrNoPath
	}
	return os.Chdir(args[1])
}

// exitShell exits the shell
func exitShell() {
	os.Exit(0)
}

// printHelp prints the help message
func printHelp() {
	fmt.Println("Supported commands:")
	fmt.Println("  dir [path]   - Show files in the directory")
	fmt.Println("  ls  [path]   - Show files in the directory")
	fmt.Println("  cd  [path]   - Change directory")
	fmt.Println("  exit         - Exit the shell")
	fmt.Println("  help         - Show this help message")
}

// handleCrossPlatformDir handles 'ls' and 'dir' commands cross-platform
func handleCrossPlatformDir(args []string) error {
	if runtime.GOOS == "windows" {
		args[0] = "dir" // Replace 'ls' with 'dir' on Windows
	} else {
		args[0] = "ls" // Use 'ls' on Unix-like systems
	}
	return executeExternalCommand(strings.Join(args, " "), args)
}

// executeExternalCommand executes external shell commands
func executeExternalCommand(input string, args []string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", append([]string{"/C"}, args...)...)
	} else {
		cmd = exec.Command("sh", "-c", input)
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
