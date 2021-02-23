package pkg

import "fmt"

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// ColorsPrint is a object that handle all color's prints
type ColorsPrint struct{}

// Cyan Println function
func (*ColorsPrint) Cyan(a string) {
	fmt.Println(colorCyan + a + colorReset)
}

// Red Println function
func (*ColorsPrint) Red(a string) {
	fmt.Println(colorRed + a + colorReset)
}

// Blue Println function
func (*ColorsPrint) Blue(a string) {
	fmt.Println(colorBlue + a + colorReset)
}

// Green Println function
func (*ColorsPrint) Green(a string) {
	fmt.Println(colorGreen + a + colorReset)
}
