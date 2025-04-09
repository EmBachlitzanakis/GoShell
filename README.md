# **GoShell**

GoShell is a lightweight, cross-platform shell built using Go. It allows you to execute basic commands, navigate directories, and interact with your system in a terminal-like environment. Designed for both Windows and Unix-like systems (Linux/macOS), GoShell provides a user-friendly interface for essential shell operations and includes dynamic features for enhanced usability.

---

## **Features**

### **Core Functionalities**
- **Cross-Platform Compatibility**: Works seamlessly on Windows, Linux, and macOS.
- **Dynamic Prompt**: Displays the current working directory as part of the shell prompt.
- **Command History**: Tracks the last 10 executed commands for quick reference.
- **Custom Command Parsing**: Distinguishes between built-in and external commands for flexible execution.
- **Command Customization**: Create custom aliases for commands stored in a JSON file, allowing you to map shorthand keywords to complex commands for personalized workflows.

---

## **Supported Commands**

### **Built-in Commands**

| Command            | Description                                      |
|--------------------|--------------------------------------------------|
| `cd [path]`        | Change the current directory.                   |
| `ls [path]`        | List files in a directory. Works cross-platform. |
| `dir [path]`       | Alias for `ls`. On Windows, uses native behavior.|
| `exit`             | Exit the shell.                                 |
| `help`             | Display the help message with available commands.|
| `history`          | Display the history of the last 10 commands.    |
| `customize "X" "Y"`| Customize a command: map alias `X` to run `Y`.   |

### **External Commands**
- If a command is not recognized as built-in, GoShell will attempt to execute it as an external system command using the native shell (`cmd` on Windows, `sh` on Unix-like systems).

---

## **New Feature: Command Customization**
GoShell introduces the ability to customize commands and store them in a JSON file for persistent usage. This feature enables you to map shorthand commands to more complex or frequently used commands. 

### **Usage Example**
1. Customize a command by entering:
   ```bash
   customize "internet" "start chrome"
This stores the alias internet in the JSON file to run the command start chrome
When you type internet. The shell automatically executes start chrome command.


### **Launch GoShell**
1. Navigate to the `GoShell` directory in your terminal.
2. Run the shell using:
   ```bash
   go run ./


