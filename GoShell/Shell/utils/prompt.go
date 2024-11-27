package utils

import (
	"fmt"
	"os"
)

func GetPrompt() string {
	dir, err := os.Getwd()
	if err != nil {
		dir = "?"
	}
	return fmt.Sprintf("%s> ", dir)
}
