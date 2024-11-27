package shell

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "Shell/commands"
    "Shell/utils"
)

func StartShell() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print(utils.GetPrompt())

        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error reading input:", err)
            continue
        }

        if err := HandleInput(input); err != nil {
            fmt.Fprintln(os.Stderr, "Error:", err)
        }
    }
}

func HandleInput(input string) error {
    input = strings.TrimSpace(input)
    if input == "" {
        return nil
    }

    // Save to history and split into args (delegated to commands package)
    commands.SaveToHistory(input)
    args := strings.Split(input, " ")
    return commands.ExecuteCommand(args)
}
