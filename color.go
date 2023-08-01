package brush

import "fmt"

// ColorType represents a color from any set
type ColorType interface {
	ANSIColor | ExtendedANSIColor

	foreground() string
	background() string
}

// ANSIColor represents a color from the first 16 colors in the ANSI table
type ANSIColor int8

// All the different 16 different colors on the ANSIColor table
const (
	Black ANSIColor = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	BrightBlack
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

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

// ToExtended transforms ANSIColor to an ExtendedANSIColor
func (color ANSIColor) ToExtended() ExtendedANSIColor {
	return ExtendedANSIColor(color)
}

// ColorIntensity rapresent a level on the scale of color color intensity from 0 to 5
type ColorIntensity uint8

const (
	ZeroIntensity     ColorIntensity = iota // Intensity lv. 0 = 0
	LowIntensity                            // Intensity lv. 1 = 95
	ModerateIntensity                       // Intensity lv. 2 = 135
	MediumIntensity                         // Intensity lv. 3 = 175
	HightIntensity                          // Intensity lv. 4 = 215
	MaxIntensity                            // Intensity lv. 5 = 255
)

// RGB picks a color from the ExtendedANSIColor table by mixing for each
// one of the primary colors, different levels of intensity from 0 to 5.
//
// A set of contant is declared on the package as helper, in order:
// ZeroIntensity, LowIntensity, ModerateIntensity, MediumIntensity, HightIntensity, MaxIntensity
func RGB(red, green, blue ColorIntensity) ExtendedANSIColor {
	return ExtendedANSIColor(16 + (36 * red) + (6 * green) + blue)
}

// MaxGrayScale represents the max value you can pass to the GrayScale
const MaxGrayScale uint8 = 25

// GrayScale pick a color from the ExtendedANSIColor table on the grayscale (1-24)
// plus the values of Black (0) and White (25)
func GrayScale(grayScale uint8) ExtendedANSIColor {
	switch grayScale {
	case 0:
		return 0
	case MaxGrayScale:
		return BrightWhite.ToExtended()
	}
	return ExtendedANSIColor(231 + grayScale)
}

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

const (
	esc        = '\x1b'
	csi        = string(esc) + "["
	colorReset = "0"
)

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
