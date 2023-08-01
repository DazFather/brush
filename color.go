package brush

import "fmt"

const (
	esc        = '\x1b'
	csi        = string(esc) + "["
	colorReset = "0"
)

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

type style struct {
	foreground, background string
}

func serialize[color ColorType](foreground color, background Optional[color]) style {
	var s = style{foreground: foreground.foreground()}
	if background != nil {
		s.background = (*background).background()
	}

	return s
}

func (b *Brush[color]) extract() style {
	return serialize(b.Foreground, b.Background)
}

func (s style) apply(content string) string {
	style := s.foreground
	if len(s.background) > 0 {
		style += ";" + s.background
	}

	return fmt.Sprintf("%s%sm%s%sm", csi, style, content, csi+colorReset)
}
