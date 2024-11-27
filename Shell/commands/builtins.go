package commands

import (
	"Shell/execution"
	"Shell/utils"
	"fmt"
	"os"
)

var commandHistory []string

const historyLimit = 10

func SaveToHistory(cmd string) {
	if len(commandHistory) >= historyLimit {
		commandHistory = commandHistory[1:]
	}
	commandHistory = append(commandHistory, cmd)
}

func ExecuteCommand(args []string) error {
	command := args[0]
	switch command {
	case "cd":
		return changeDirectory(args)
	case "history":
		printHistory()
		return nil
	case "help":
		printHelp()
		return nil
	case "ls", "dir":
		// Unified handling for directory listing
		return listDirectory(args)
	default:
		// Pass unrecognized commands to the external execution logic
		return execution.ExecuteExternalCommand(args)
	}
}

func changeDirectory(args []string) error {
	if len(args) < 2 {
		return utils.ErrNoPath
	}
	return os.Chdir(args[1])
}

func listDirectory(args []string) error {
	// Default to current directory if no path is provided
	dir := "."
	if len(args) > 1 {
		dir = args[1]
	}

	// Open the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory '%s': %v", dir, err)
	}

	// Print directory contents
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("[DIR]  %s\n", file.Name())
		} else {
			fmt.Printf("       %s\n", file.Name())
		}
	}
	return nil
}

func printHistory() {
	for i, cmd := range commandHistory {
		fmt.Printf("%d: %s\n", i+1, cmd)
	}
}

func printHelp() {
	fmt.Println("Supported commands:")
	fmt.Println("  cd [path]    - Change directory")
	fmt.Println("  history      - Show command history")
	fmt.Println("  help         - Show help message")
	fmt.Println("  dir [path]   - Show files in the directory")
	fmt.Println("  ls  [path]   - Show files in the directory")
	fmt.Println("  [command]    - Execute external shell command")
}
