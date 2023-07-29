package brush

import "fmt"

type ColorType interface {
	ANSIColor | ExtendedANSIColor

	foreground() string
	background() string
}

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

type ExtendedANSIColor uint8

func (c ExtendedANSIColor) foreground() string {
	return fmt.Sprint("38;5;", int(c))
}

func (c ExtendedANSIColor) background() string {
	return fmt.Sprint("48;5;", int(c))
}

type Optional[color ColorType] *color

func UseColor[color ColorType](c color) Optional[color] {
	return &c
}

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
