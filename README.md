<p align="center" width="100%">
	<a href="https://github.com/DazFather/brush/raw/main/examples/logo/main.go">
		<img alt="brush logo" src="https://github.com/DazFather/brush/raw/main/examples/logo/github_dazfather_brush_logo.png">
	</a>
	<p align="center" width="100%">
		<a href="https://img.shields.io/badge/Language-Go-blue.svg"><img alt="Language" src="https://img.shields.io/badge/Language-Go-blue.svg"></a>
		<a href="https://github.com/DazFather/brush/blob/main/LICENSE"><img alt="License" src="http://img.shields.io/badge/license-MIT-orange.svg?style=flat"></a>
		<a href="https://goreportcard.com/report/github.com/DazFather/brush"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/DazFather/brush"></a>
		<a href="https://coveralls.io/github/DazFather/brush?branch=main"><img alt="Coverage Status" src="https://coveralls.io/repos/github/DazFather/brush/badge.svg?branch=main"/></a>
		<a href="https://pkg.go.dev/github.com/DazFather/brush"><img alt="Go Reference" src="https://pkg.go.dev/badge/github.com/DazFather/brush.svg"></a>
	</p>
</p>

Brush is a simple and light library to help you paint your terminal outputs

>  Heavily inspired by [termenv](https://github.com/muesli/termenv) and [lipgloss](https://github.com/charmbracelet/lipgloss)

## Philosopy

 - **Simple** and **Lightway** - _it's a brush not a space shuttle!_
 > _"Less is exponentially better"_ I do believe in it and I tried to apply it on this library,
 > no many complex data structures or even just functions/methods, no external libraries.
 - **Safe** - _16 != 256_
 > When handling special char sequence is very common to commit mistakes, for this reason
 > when you paint something you don't have a string but a `Painted` item with a `String() string` method
 > (The same applied to `Highlighted` items).
 > Another very common problem, when handling ANSI color codes, is mixing the original 16 palette with
 > the extended (256) version. When creating a brush or painting / hightlighting something
 > font and background colors must be of the same ColorType   
 - **Idiomatic** - _mybrush := brush.New(...)_
 > Naming convention helps to make the library easy to use and intuitive
 > throw the use of functions/methods like Paint, Paintln, Paintf to color stuffs
 > that acts very similarly to the one defined on the fmt package
 

## Usage

**Paint** Use the [Paint](https://pkg.go.dev/github.com/DazFather/brush#Paint) function to "paint" a given string with a foreground and background color, if the background is `nil` then the background will be the default of your terminal.
```go
fmt.Println("I", brush.Paint(brush.Red, nil, "love"), "go")
```
Create your own brush via the [New](https://pkg.go.dev/github.com/DazFather/brush#New) function so you don't have to pass the style every time you want to paint something.
```go
myBrush := brush.New(brush.Red, nil)
fmt.Println("I", myBrush.Paint("love"), "go")

// You can use different methods to change the style of the brush like UseFontColor, UseBgColor
myBrush.UseBgColor(brush.Black).UseFontColor(brush.Yellow) // you can chain them!

fmt.Println(myBrush.Paint("Hello"), myBrush.Swap().Paint("World"), "!") // Swap will invert font and bg color
fmt.Println(
	myBrush.UseBgTransparent().Paint("I"), // UseBgTransparent will remove the bg
	myBrush.UseDefaultColor().Paint("love"), // UseDefaultColor will reset the colors to the ones on brush declaration
	"go",
)
```
Use the [Highlight](https://pkg.go.dev/github.com/DazFather/brush#Highlight) method to color just the matching part of the string
```go
fmt.Println(myBrush.Hightlight("I love go", regexp.MustCompile("love")))
```

### Examples
If you need more examples, you can find more [here](https://github.com/DazFather/brush/tree/main/examples) 


## Colors
The library uses the ANSI Color codes in 3 different format:
 > I like to keep things separated for safety so when you declare a new brush or painting something be sure to use both colors (font and background) from the same `ColorType`.

### ANSIColor
For the original ANSI (16 colors) you can simply use the already defined constants:
 > `Black`, `Red`, `Green`, `Yellow`, `Blue`, `Magenta`, `Cyan`, `White` <br>
 > and their "bright" version: <br>
 > `BrightBlack`, `BrightRed`, `BrightGreen`, `BrightYellow`, `BrightBlue`, `BrightMagenta`, `BrightCyan`, `BrightWhite`

Or use a simple interger (`int8`)

### ExtendedANSIColor
For the extended ANSI (256 colors) you can or use one of the first 16 colors by using the method: `ToExtended()`
 > ex. `font := brush.Red.ToExtended()`

As for the ANSIColor scheme you use a simple interger (`uint8`)

For the last 24 colors (the grayscale) + black and white, you can use the the `GrayScale` function
 > White is the maximum value: `myWhite := brush.GrayScale(25)`, you can also use the constant `MaxGrayScale` <br>
 > Black is the minimum value: `myBlack := brush.GrayScale(0)` <br>
 > All the other values in between represents the grayscale (the last 24 colors of the ANSI table): `myGray := brush.GrayScale(16)` 

For the central part of the colors codes table you can compose them using their `RGB` values with the `RGB` function
 > You simply need to put the value or red, green and blue `myBlue := brush.RGB(0, 0, 5)`. <br>
 > Keep in mind throw that the scale goes from 0 to 5. To help with that there is a set of constant you can use: <br>
 > `ZeroIntensity` (0), `LowIntensity` (95), `ModerateIntensity` (135), `MediumIntensity` (175), `HightIntensity` (215), `MaxIntensity` (255) <br>
 > If you need all the possible colors with rgb values ranging from 0 to 255 then...

### TrueColor
Not supported in every terminal but it allows to use all the red, green and blue values ranging from 0 to 255. 
For the previous colors you can use the method `ToTrueColor()` provided for ANSIColor and ExtendedANSIColor ColorType
 > ex. `font := brush.Red.ToTrueColor()`

For the RGB is very easy: `TrueColor` is a struct with the Red, Green and Blue fields (`uint16`), so just initialize one
 > ex. `yellow := brush.TrueColor{255, 165, 0}`

If you need to convert an **hexadecimal** color, you can use the `ParseHex` function, it accept both `#RRGGBB` and `#RGB` formats,
the `#` is totally optional 
 > ex. `yellowPtr, err := brush.ParseHex("FFA500")`

