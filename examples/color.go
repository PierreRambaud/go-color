package main

import(
	"github.com/PierreRambaud/go-color"
	"strings"
)

func main() {
	color.New("fgcyan", "bgcyan", "underline").Println("something")
	color.New("fgred", "bold").Println("something")

	for fgk, _ := range color.FgColors {
		for bgk, _ := range color.BgColors {
			color.New(fgk, strings.Join([]string{"bg", bgk}, "")).Print(strings.Join([]string{fgk, bgk}, " "))
			println("")
		}
	}
}
