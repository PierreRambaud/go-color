package color

import(
	"fmt"
	"bytes"
	"testing"
)

func TestColor(t *testing.T) {
	rb := new(bytes.Buffer)
	colorOutput = rb

	testColors := []struct {
		text string
		code int
		style string
	}{
		{text: "foreground black", style: "fgblack", code: FgColors["black"]},
		{text: "foreground red", style: "fgred", code: FgColors["red"]},
		{text: "foreground green", style: "fggreen", code: FgColors["green"]},
		{text: "foreground yellow", style: "fgyellow", code: FgColors["yellow"]},
		{text: "foreground blue", style: "fgblue", code: FgColors["blue"]},
		{text: "foreground magenta", style: "fgmagenta", code: FgColors["magenta"]},
		{text: "foreground white", style: "fgwhite", code: FgColors["white"]},
		{text: "background black", style: "bgblack", code: BgColors["black"]},
		{text: "background red", style: "bgred", code: BgColors["red"]},
		{text: "background green", style: "bggreen", code: BgColors["green"]},
		{text: "background yellow", style: "bgyellow", code: BgColors["yellow"]},
		{text: "background blue", style: "bgblue", code: BgColors["blue"]},
		{text: "background magenta", style: "bgmagenta", code: BgColors["magenta"]},
		{text: "background white", style: "bgwhite", code: BgColors["white"]},
		{text: "bold", style: "bold", code: Style["bold"]},
		{text: "faint", style: "faint", code: Style["faint"]},
		{text: "italic", style: "italic", code: Style["italic"]},
		{text: "underline", style: "underline", code: Style["underline"]},
		{text: "blink slow", style: "blinkslow", code: Style["blinkslow"]},
		{text: "blink rapid", style: "blinkrapid", code: Style["blinkrapid"]},
		{text: "reverse", style: "reverse", code: Style["reverse"]},
		{text: "conceal", style: "conceal", code: Style["conceal"]},
		{text: "crossed out", style: "crossedout", code: Style["crossedout"]},
	}

	for _, c := range testColors {
		New(c.style).Print(c.text)
		test_return(t, rb, "\x1b[%dm%s\x1b[0m", c.code, c.text)
	}

	New("red", "underline").Print("red underline")
	test_return(t, rb, "\x1b[%d;4m%s\x1b[0m", FgColors["red"], "red underline")

	New("nothing").Print("nothing")
	test_return(t, rb, "\x1b[%dm%s\x1b[0m", 0, "nothing")

	success := New("green").Func("Print")
	success("Oh YEAH!")
	test_return(t, rb, "\x1b[%dm%s\x1b[0m", FgColors["green"], "Oh YEAH!")
}

func test_return(t *testing.T, rb *bytes.Buffer, format string, code int, text string) bool {
	line, _ := rb.ReadString('\n')
	returnLine := fmt.Sprintf("%q", line)
	colored := fmt.Sprintf(format, code, text)
	expectLine := fmt.Sprintf("%q", colored)
	fmt.Printf("%s  \t: %s\n", text, line)
	if returnLine != expectLine {
		t.Errorf("Expecting %s, got '%s'\n", expectLine, returnLine)
	}

	return true
}
