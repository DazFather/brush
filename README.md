# Brush
[![Go Report Card](https://goreportcard.com/badge/github.com/DazFather/brush)](https://goreportcard.com/report/github.com/DazFather/brush)
[![Go Reference](https://pkg.go.dev/badge/github.com/DazFather/brush.svg)](https://pkg.go.dev/github.com/DazFather/brush)
> Your humble terminal paintbrush. Heavily inspired by [termenv](https://github.com/muesli/termenv) and [lipgloss](https://github.com/charmbracelet/lipgloss)


Brush is a simple and light library to help you paint your terminal outputs. If you need more than this consider using one of the libraries linked before.

## Usage

Use the default "brush" aka the [Paint](https://pkg.go.dev/github.com/DazFather/brush#Paint) function to "paint" a given string with a foreground and background color, if the background is nil then the background will be the default of your terminal.
```go
package main

import (
	"fmt"

	"github.com/DazFather/brush"
)

func main() {
	// The nil indicates the background (in this case Transparent)
	fmt.Println("I", brush.Paint(brush.Red, nil, "love"), "go")
}
```
Create your own brush via the [New](https://pkg.go.dev/github.com/DazFather/brush#New) function to change between multiple styling
```go
package main

import (
	"fmt"

	"github.com/DazFather/brush"
)

func main() {
	myBrush := brush.New(brush.Yellow, brush.UseColor(brush.Black))

	fmt.Println(myBrush.Paint("Hello"), myBrush.Swap().Paint("World"), "!")
	fmt.Println(myBrush.UseDefaultColor().Paint("I"), myBrush.UseFontColor(brush.Red).Paint("love"), "go")
}
```

## Colors
The library uses the ANSI Color codes in 2 different format: `ANSIColor` (16 colors) and `ExtendedANSIColor` (256 colors).
Check out the color codes on this table:
![ANSI color chart](https://github.com/muesli/termenv/raw/master/examples/color-chart/color-chart.png)
> Thanks again [termenv](https://github.com/muesli/termenv) for the image, don't mind if I _yoink_ it for now, right? <3

## To Do
- [x] Create a function to embed a painted item into a string to paint
- [x] Add documentation
- [x] Add some helper functions and constants for selecting colors
- [ ] Create some examples and add some screenshots in the README

