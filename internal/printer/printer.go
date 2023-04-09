package printer

import (
	"fmt"

	"github.com/muesli/termenv"
)

func Print(s string, stringType string) {
	p := termenv.ColorProfile()
	styled := termenv.String(s)
	switch stringType {
	case "welcome":
		fmt.Println(styled.Foreground(p.Color("#FF2FF2")))

	case "theme":
		fmt.Println(styled.Foreground(p.Color("#FF3E19")))

	case "error":
		fmt.Println(styled.Foreground(p.Color("#FF0000")))

	case "tip":
		fmt.Println(styled.Foreground(p.Color("#00FFF0")))

	case "guide":
		fmt.Println(styled.Foreground(p.Color("#00FF00")))

	default:
		fmt.Println(s)
	}
}
