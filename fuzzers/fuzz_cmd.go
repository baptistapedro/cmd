package myfuzz

import "github.com/commander-cli/cmd"

// CreateNewCommandWithStandardStream create new standard stream example
func Fuzz(data []byte) int {
	c := cmd.NewCommand(string(data), cmd.WithStandardStreams)
	c.Execute()
	return 0
}
