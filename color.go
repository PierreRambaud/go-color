package color

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Define color object
type Color struct {
	params []int
}

const Escape = "\x1b"

const(
	Reset = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

const(
	FgBlack int = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

const(
	BgBlack int = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

var output io.Writer = os.Stdout

func New(value ...int) *Color {
	c := &Color{params: make([]int, 0)}
	c.add(value...)
	return c
}

func (c *Color) add(value ...int) *Color {
	c.params = append(c.params, value...)
	return c
}

func (c *Color) set() *Color {
	fmt.Fprintf(output, c.format())
	return c
}

func (c *Color) unset() *Color {
	fmt.Fprintf(output, "%s[%dm", Escape, Reset)
	return c
}

func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for k, v := range c.params {
		format[k] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", Escape, c.sequence())
}

func (c *Color) unformat() string {
	return fmt.Sprintf("%s[%dm", Escape, Reset)
}

func (c *Color) Print(attr ...interface{}) (n int, err error) {
	c.set()
	defer c.unset()
	return fmt.Fprint(output, attr...)
}

func (c *Color) Printf(format string, attr... interface{}) (n int, err error) {
	c.set()
	defer c.unset()
	return fmt.Fprintf(output, format, attr...)
}

func (c *Color) Println(attr ...interface{}) (n int, err error) {
	c.set()
	defer c.unset()
	return fmt.Fprintln(output, attr...)
}
