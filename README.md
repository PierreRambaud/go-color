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
success := color.New("green").add("bold")
success.Print("Great!")
success.Printf("This is %s", "SPARTA")
```

## TODO

- Reusable functions


## License

see [LICENSE.md](LICENSE.md) for more details
