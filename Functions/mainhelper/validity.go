package mainhelper

import (
	"fmt"
	"reflect"

	"netcat/Functions/natheerspretty"
)

func IsEmpty(message []byte) bool {
	// for _, v := range message {
	// 	fmt.Println(natheerspretty.RGBify(255, 0, 255, v))
	// }

	if len(message) <= 0 {
		fmt.Println(natheerspretty.RGBify(255, 0, 0, "the message is empty"))
		return true
	}

	for i := 0; i < len(message); i++ {
		// check if the message is all spaces
		if message[i] > 32 {
			return false
		}
	}
	fmt.Println(natheerspretty.RGBify(255, 0, 0, "the message is full of spaces"))

	return true
}

func Signaltrapchecker(message []byte) bool {
	UP := []byte{27, 91, 65}
	DOWN := []byte{27, 91, 66}
	LEFT := []byte{27, 91, 68}
	RIGHT := []byte{27, 91, 67}
	HOME := []byte{27, 91, 72}
	END := []byte{27, 91, 70}
	sigtraps := [][]byte{UP, DOWN, LEFT, RIGHT, HOME, END}

	if len(message) > 2 {
		for i := 0; i < len(message)-3+1; i++ {
			for _, sigtrap := range sigtraps {
				if reflect.DeepEqual(message[i:i+3], sigtrap) {
					// fmt.Println("detected", sigtrap)
					return true
				}
			}
		}
	}
	return false
}
