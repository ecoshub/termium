package ansi

import "fmt"

func PrintAllColors() {
	for i := 0; i < 256; i++ {
		// for j := 0; j < 256; j++ {
		// }
		if i%8 == 0 {
			print("\n")
		}
		fmt.Printf("\x1b[38;5;255m\x1b[48;5;%dm  [%3d]  \x1b[0m ", i, i)
	}
	print("\n")
}

func ResetAllModes() string {
	return EscapeChar + "[0m"
}

func ColorGray() string {
	return EscapeChar + "[1;66m"
}

func ColorRed() string {
	return EscapeChar + "[1;66m"
}
