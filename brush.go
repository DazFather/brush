package brush

import (
	"fmt"
	"strings"
)

type Brush[color ColorType] struct {
	Foreground, defForeground color
	Background, defBackground Optional[color]
}

func New[color ColorType](font color, background Optional[color]) Brush[color] {
	var b = Brush[color]{
		defForeground: font,
		defBackground: background,
	}
	b.UseDefaultColor()

	return b
}

func (b *Brush[color]) UseFontColor(c color) *Brush[color] {
	b.Foreground = c
	return b
}

func (b *Brush[color]) UseBgColor(c color) *Brush[color] {
	b.Background = UseColor(c)
	return b
}

func (b *Brush[color]) UseBgTransparent() *Brush[color] {
	b.Background = nil
	return b
}

func (b *Brush[color]) Swap() *Brush[color] {
	b.Background, b.Foreground = UseColor(b.Foreground), PickColor(b.Background, b.defForeground)
	return b
}

func (b *Brush[color]) UseDefaultColor() *Brush[color] {
	b.Foreground, b.Background = b.defForeground, b.defBackground
	return b
}

func (b Brush[color]) Paint(s ...string) painted {
	return painted{
		foreground: b.Foreground.foreground(),
		background: serializeBg(b.Background),
		origin:     strings.Join(s, ""),
	}
}

func (b Brush[color]) Repaint(p ...painted) painted {
	var result = painted{
		foreground: b.Foreground.foreground(),
		background: serializeBg(b.Background),
	}

	for i := range p {
		result.origin += p[i].origin
	}

	return result
}

func (b Brush[color]) Print(s ...string) {
	fmt.Print(b.Paint(s...))
}

func (b Brush[color]) Println(s ...string) {
	b.Print(strings.Join(s, " "), "\n")
}
