package mainhelper

import (
	"fmt"
	"netcat/Functions/natheerspretty"
)

func IsEmpty(message []byte) bool {
	if len(message) <= 0 {
		fmt.Println(natheerspretty.RGBify(255, 0, 0, "the message is empty"))
		return true
	}

	for i := 0; i < len(message); i++ {
		//check if the message is all spaces
		if message[i] != 32 {
			return false
		}
	}
	fmt.Println(natheerspretty.RGBify(255, 0, 0, "the message is full of spaces"))

	return true
}
