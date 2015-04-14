package color

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"reflect"
	"regexp"
)

// Define color object
type Color struct {
	params []int
}

const Escape = "\x1b"

const(
	black int = iota
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

var colorOutput io.Writer = os.Stdout

var BgColors = map[string]int{
	"black": black + 40,
	"red": red + 40,
	"green": green + 40,
	"yellow": yellow + 40,
	"blue": blue + 40,
	"magenta": magenta + 40,
	"cyan": cyan + 40,
	"white": white + 40,
}

var FgColors = map[string]int{
	"black": black + 30,
	"red": red + 30,
	"green": green + 30,
	"yellow": yellow + 30,
	"blue": blue + 30,
	"magenta": magenta + 30,
	"cyan": cyan + 30,
	"white": white + 30,
}

var Style = map[string]int {
	"reset": 0,
	"bold": 1,
	"faint": 2,
	"italic": 3,
	"underline": 4,
	"blinkslow": 5,
	"blinkrapid": 6,
	"reverse": 7,
	"conceal": 8,
	"crossedout": 9,
}

func New(value ...string) *Color {
	c := &Color{params: make([]int, 0)}
	for _, v := range value {
		c.Add(v)
	}

	return c
}

func (c *Color) Add(value ...string) *Color {
	for _, v := range value {
		c.params = append(c.params, ColorCode(v))
	}

	return c
}

func (c *Color) Set() *Color {
	fmt.Fprintf(colorOutput, c.Format())
	return c
}

func (c *Color) Unset() *Color {
	fmt.Fprintf(colorOutput, "%s[%dm", Escape, Style["reset"])
	return c
}

func (c *Color) Sequence() string {
	format := make([]string, len(c.params))
	for k, v := range c.params {
		format[k] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

func (c *Color) Format() string {
	return fmt.Sprintf("%s[%sm", Escape, c.Sequence())
}

func (c *Color) Print(attr ...interface{}) (n int, err error) {
	c.Set()
	defer c.Unset()
	return fmt.Fprint(colorOutput, attr...)
}

func (c *Color) Printf(format string, attr... interface{}) (n int, err error) {
	c.Set()
	defer c.Unset()
	return fmt.Fprintf(colorOutput, format, attr...)
}

func (c *Color) Println(attr ...interface{}) (n int, err error) {
	c.Set()
	defer c.Unset()
	return fmt.Fprintln(colorOutput, attr...)
}

func (c *Color) Func(method string) func(args ...interface{}) {
	return func(args ...interface{}) {
		inputs := make([]reflect.Value, len(args))
		for i, _ := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}

		reflect.ValueOf(c).MethodByName(method).Call(inputs)
	}
}

func ColorCode(code string) int {
	if val, ok := Style[code]; ok {
		return val
	}

	if color := MatchString("^fg(.*)", code); color != "" {
		return FgColors[color]
	}

	if val, ok := FgColors[code]; ok {
		return val
	}

	if color := MatchString("^bg(.*)", code); color != "" {
		return BgColors[color]
	}

	return 0
}

func MatchString(regex string, code string) string {
	r, _ := regexp.Compile(regex)
	match := r.FindStringSubmatch(code)
	if len(match) != 0 {
		return match[1]
	}

	return ""
}
