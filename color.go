package brush

import "github.com/muesli/termenv"

type ColorType interface {
	termenv.ANSI256Color | termenv.ANSIColor | termenv.RGBColor

	Sequence(bool) string
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

func sequenceBg[color ColorType](opt Optional[color]) string {
	if opt == nil {
		return ""
	}

	return (*opt).Sequence(true)
}
