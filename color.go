package brush

import "fmt"

// ColorType represents a color from any set
type ColorType interface {
	ANSIColor | ExtendedANSIColor

	foreground() string
	background() string
}

// ANSIColor represents a color from the first 16 colors in the ANSI table
type ANSIColor uint8

func (c ANSIColor) foreground() string {
	col := int(c)

	if col < 8 {
		return fmt.Sprint(col + 30)
	}
	return fmt.Sprint(col + 82)
}

func (c ANSIColor) background() string {
	col := int(c)

	if col < 8 {
		return fmt.Sprint(col + 40)
	}
	return fmt.Sprint(col + 92)
}

// ExtendedANSIColor represents a color from the extended ANSI table (256 colors)
type ExtendedANSIColor uint8

func (c ExtendedANSIColor) foreground() string {
	return fmt.Sprint("38;5;", int(c))
}

func (c ExtendedANSIColor) background() string {
	return fmt.Sprint("48;5;", int(c))
}

// Optional represents an optional color
type Optional[color ColorType] *color

// UseColor is an utility that lets you transform a color in an Optional
// this can be expecially usefull on the New or Paint function when selecting background
func UseColor[color ColorType](c color) Optional[color] {
	return &c
}

// PickColor is an utility that lets you take the color referenced on the first argument
// or in case it's nil the second as a default case
func PickColor[color ColorType](opt Optional[color], def color) color {
	if opt != nil {
		return *opt
	}

	return def
}

func serializeBg[color ColorType](opt Optional[color]) (sequence string) {
	if opt != nil {
		sequence = (*opt).background()
	}

	return
}
