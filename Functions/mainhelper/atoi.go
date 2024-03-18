package mainhelper

import (
	"errors"
)

func Atoi(s string) (int, error) {
	arr := []rune(s)
	converted := 0
	multi := 1
	sign := 0
	for i := 0; i < len(arr)-1; i++ {
		if (arr[0] > 57) || (arr[0] < 48) && (arr[0] != 45) && (arr[0] != 43) {
			return 0, errors.New("invalid first character: the first character must be a digit, minus sign (-), or plus sign (+)")
		}
		if arr[i+1] > 57 || arr[i+1] < 48 {
			return 0, errors.New("invalid character: only digits are allowed")
		}
		if arr[0] == 45 {
			sign = -1
			arr = arr[1:]
		}
		if arr[0] == 43 {
			sign = 1
			arr = arr[1:]
		}
	}
	for i := 0; i < len(arr)-1; i++ {
		multi = multi * 10
	}
	if multi == 1 {
		if len(arr) == 0 {
			return 0, errors.New("empty input: the input string is empty")
		}
		if arr[0] == 45 {
			return 0, errors.New("invalid input: negative sign without digits")
		}
		if arr[0] == 43 {
			return 0, errors.New("invalid input: positive sign without digits")
		}
	}
	if sign == 0 {
		for i := 0; i <= len(arr)-1; i++ {
			if arr[i] < 48 || arr[i] > 57 {
				return 0, errors.New("invalid character: only digits are allowed")
			}
			if converted > (1<<63-1)/10 || (converted == (1<<63-1)/10 && int(arr[i])-'0' > 7) {
				return 0, errors.New("overflow: the value exceeds the maximum representable integer")
			}
			converted = converted*10 + (int(arr[i]) - '0')
		}
	}
	if sign == -1 {
		for i := 0; i <= len(arr)-1; i++ {
			if arr[i] < 48 || arr[i] > 57 {
				return 0, errors.New("invalid character: only digits are allowed")
			}
			if converted > (1<<63-1)/10 || (converted == (1<<63-1)/10 && int(arr[i])-'0' > 8) {
				return 0, errors.New("underflow: the value is less than the minimum representable integer")
			}
			converted = converted*10 - (int(arr[i]) - '0')
		}
	}
	if sign == 1 {
		for i := 0; i <= len(arr)-1; i++ {
			if arr[i] < 48 || arr[i] > 57 {
				return 0, errors.New("invalid character: only digits are allowed")
			}
			if converted > (1<<63-1)/10 || (converted == (1<<63-1)/10 && int(arr[i])-'0' > 7) {
				return 0, errors.New("overflow: the value exceeds the maximum representable integer")
			}
			converted = converted*10 + (int(arr[i]) - '0')
		}
	}
	return converted, nil
}
