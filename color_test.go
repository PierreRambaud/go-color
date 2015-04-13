package color

import(
	"fmt"
	"bytes"
	"testing"
)

func TestColor(t *testing.T) {
	rb := new(bytes.Buffer)
	output = rb

	testColors := []struct {
		text string
		code int
	}{
		{text: "foreground black", code: FgBlack},
		{text: "foreground red", code: FgRed},
		{text: "foreground green", code: FgGreen},
		{text: "foreground yellow", code: FgYellow},
		{text: "foreground blue", code: FgBlue},
		{text: "foreground magenta", code: FgMagenta},
		{text: "foreground white", code: FgWhite},
		{text: "background black", code: BgBlack},
		{text: "background red", code: BgRed},
		{text: "background green", code: BgGreen},
		{text: "background yellow", code: BgYellow},
		{text: "background blue", code: BgBlue},
		{text: "background magenta", code: BgMagenta},
		{text: "background white", code: BgWhite},
		{text: "bold", code: Bold},
		{text: "faint", code: Faint},
		{text: "italic", code: Italic},
		{text: "underline", code: Underline},
		{text: "blink slow", code: BlinkSlow},
		{text: "blink rapid", code: BlinkRapid},
		{text: "reverse video", code: ReverseVideo},
		{text: "concealed", code: Concealed},
		{text: "crossed out", code: CrossedOut},
	}

	for _, c := range testColors {
		New(c.code).Print(c.text)

		line, _ := rb.ReadString('\n')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", c.code, c.text)
		expectLine := fmt.Sprintf("%q", colored)
		fmt.Printf("%s\t: %s\n", c.text, line)
		if scannedLine != expectLine {
			t.Errorf("Expecting %s, got '%s'\n", expectLine, scannedLine)
		}
	}
}
