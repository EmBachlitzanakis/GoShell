package commands

import (
	"Shell/execution"
	"Shell/utils"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var commandHistory []string

const aliasFile = "command_aliases.json"
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
		return listDirectory(args)
	case "customize":
		return customizeCommand(args)
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
	dir := "."
	if len(args) > 1 {
		dir = args[1]
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory '%s': %v", dir, err)
	}

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

func customizeCommand(args []string) error {

	input := strings.Join(args, " ")

	pattern := `^customize\s+"(.*?)"\s+"(.*?)"$`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(input)
	if len(matches) != 3 {
		return fmt.Errorf("usage: customize \"old command\" \"new command\"")
	}

	oldCommand := matches[1]
	newCommand := matches[2]

	aliases, err := loadAliases()
	if err != nil {
		return fmt.Errorf("failed to load aliases: %v", err)
	}

	aliases[oldCommand] = newCommand

	err = saveAliases(aliases)
	if err != nil {
		return fmt.Errorf("failed to save aliases: %v", err)
	}

	fmt.Printf("Command '%s' has been customized to '%s'\n", oldCommand, newCommand)
	return nil
}

// loadAliases reads the existing aliases from the JSON file
func loadAliases() (map[string]string, error) {
	aliases := make(map[string]string)

	if _, err := os.Stat(aliasFile); os.IsNotExist(err) {
		return aliases, nil
	}

	data, err := os.ReadFile(aliasFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &aliases)
	if err != nil {
		return nil, err
	}

	return aliases, nil
}

// saveAliases writes the aliases to the JSON file
func saveAliases(aliases map[string]string) error {

	data, err := json.MarshalIndent(aliases, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(aliasFile, data, 0644)
}
