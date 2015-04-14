# Go Color
[![Build Status](https://travis-ci.org/PierreRambaud/go-color.svg?branch=master)](https://travis-ci.org/PierreRambaud/go-color)

Color your console in Golang.

## Install

```
go install github.com/PierreRambaud/go-color
```

## Examples

```go
// Default usage
color.New("green").Print("This is a green message.")
color.New("red", "underline").Print("This is a red message with an underline.")

// Reuse color
success := color.New("green").Add("bold")
success.Print("Great!")
success.Printf("This is %s", "SPARTA")


// Reusable functions
// Func expects as parameter to be a valid method from
// the Color struct.

error := color.New("red").Func("Print")
error("Error during process..")
error("Can connect..")

info := color.New("yellow").Func("Println")
info("Be careful")
info("Something happened, but nobody cares.")
```

## License

see [LICENSE.md](LICENSE.md) for more details
