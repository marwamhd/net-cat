package natheerspretty

import "fmt"

const (
	yellowANSI = "\033[33m"
	redANSI    = "\033[31m"
	greenANSI  = "\033[32m"
	resetANSI  = "\033[0m"
)

func Yellowify(args ...interface{}) string {
	return fmt.Sprint(yellowANSI, fmt.Sprint(args...), resetANSI)
}

func Redify(args ...interface{}) string {
	return fmt.Sprint(redANSI, fmt.Sprint(args...), resetANSI)
}

func Greenify(args ...interface{}) string {
	return fmt.Sprint(greenANSI, fmt.Sprint(args...), resetANSI)
}

func RGBify(r, g, b int, args ...interface{}) string {
	// Clamp RGB values between 0 and 255
	r = outofrange(r, 0, 255)
	g = outofrange(g, 0, 255)
	b = outofrange(b, 0, 255)

	ansiValue := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	return fmt.Sprint(ansiValue, fmt.Sprint(args...), resetANSI)
}

func outofrange(value, min, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}
