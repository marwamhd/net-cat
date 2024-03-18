package mainhelper

import (
	"fmt"
)

// given a number of arguments, it will return an error if the number of arguments is not the expected and the number of arguments
func ArgsValidator(args []string, expected int) (int, error) {
	if len(args) != expected {
		return 0, fmt.Errorf("error: too many arguments")
	}
	return len(args), nil
}
