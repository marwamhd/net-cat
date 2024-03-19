package mainhelper

import (
	"fmt"

	"netcat/Functions/natheerspretty"
)

func IsEmpty(message []byte) bool {
	for _, v := range message {
		fmt.Println(natheerspretty.RGBify(255, 0, 255, v))
	}

	// sigtraps:= []byte{}

	if len(message) > 2 {
		for i := 0; i < len(message)-3+1; i++ {
			fmt.Println(message[i : i+3])
			if string(message[i:i+3]) == string([]byte{27, 91, 65}) {
				fmt.Println("detected this")
			}
			// message[0:3]== [27 91 65]
			// message[1:4]== [27 91 65]
		}
	}
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
