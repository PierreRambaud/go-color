package main

import(
    "github.com/PierreRambaud/go-color"
    "strings"
)

func main() {
    // Default usage
    color.New("green").Println("This is a green message.")
    color.New("red", "underline").Println("This is a red message with an underline.")

    // Reuse color
    success := color.New("green")
    success.Add("bold")
    success.Println("Great!")
    success.Printf("This is %s", "SPARTA")
    println("")

    error := color.New("red").Func("Print")
    error("Error during process..")
    error("Can connect..")
	println("")

    info := color.New("yellow").Func("Println")
    info("Be careful")
    info("Something happened, but nobody cares.")

    for fgk, _ := range color.FgColors {
        for bgk, _ := range color.BgColors {
            color.New(fgk, strings.Join([]string{"bg", bgk}, "")).Print(strings.Join([]string{fgk, bgk}, " "))
            println("")
        }
    }
}
