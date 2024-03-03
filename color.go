package brush

import "fmt"

// ColorType represents a color from any set
type ColorType interface {
	ANSIColor | ExtendedANSIColor | TrueColor

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

// ToExtended transforms an ANSIColor to an ExtendedANSIColor representation
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
// If you want all range from 0 to 255 use TrueColor (might not be supported in your terminal)
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
// this can be especially useful on the New or Paint function when selecting a background
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

// TrueColor is a true RGB color representation.
// Be aware that not all terminal support this format
type TrueColor struct {
	Red, Green, Blue uint8
}

func (c TrueColor) foreground() string {
	return fmt.Sprint("38;2;", c.Red, ";", c.Green, ";", c.Blue)
}

func (c TrueColor) background() string {
	return fmt.Sprint("48;2;", c.Red, ";", c.Green, ";", c.Blue)
}

// ToTrueColor transforms an ANSIColor to a standard TrueColor representation.
// Be aware that the actual color might be different from the original,
// because the visible color might be different from the one of your terminal.
func (c ANSIColor) ToTrueColor() (tc TrueColor) {
	switch c {
	case Black:
		tc = TrueColor{0, 0, 0}
	case Red:
		tc = TrueColor{128, 0, 0}
	case Green:
		tc = TrueColor{0, 128, 0}
	case Yellow:
		tc = TrueColor{128, 128, 0}
	case Blue:
		tc = TrueColor{0, 0, 128}
	case Magenta:
		tc = TrueColor{128, 0, 128}
	case Cyan:
		tc = TrueColor{0, 128, 128}
	case White:
		tc = TrueColor{192, 192, 192}
	case BrightBlack:
		tc = TrueColor{128, 128, 128}
	case BrightRed:
		tc = TrueColor{255, 0, 0}
	case BrightGreen:
		tc = TrueColor{0, 255, 0}
	case BrightYellow:
		tc = TrueColor{255, 255, 0}
	case BrightBlue:
		tc = TrueColor{0, 0, 255}
	case BrightMagenta:
		tc = TrueColor{255, 0, 255}
	case BrightCyan:
		tc = TrueColor{0, 255, 255}
	case BrightWhite:
		tc = TrueColor{255, 255, 255}
	}
	return
}

// ToTrueColor transforms an ExtendedANSIColor to a TrueColor representation
func (c ExtendedANSIColor) ToTrueColor() TrueColor {
	if c < 16 {
		return ANSIColor(c).ToTrueColor()
	} else if c > 232 {
		gray := uint8((c-232)*10 + 8)
		return TrueColor{gray, gray, gray}
	}

	ansi := uint8(c - 16)
	return TrueColor{
		Blue:  ansi % 6,
		Green: (ansi / 6) % 6,
		Red:   (ansi / 36) % 6,
	}
}
