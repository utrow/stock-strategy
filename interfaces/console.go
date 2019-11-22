package interfaces

import "os"

type UserInput struct {
	HistoryFile string
}

func GetArgInputs() *UserInput {
	args := os.Args[1:]
	if len(args) < 1 {
		return nil
	}

	return &UserInput{
		HistoryFile: args[0],
	}
}
