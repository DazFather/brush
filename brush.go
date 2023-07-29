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
	return Paint(b.Foreground, b.Background, s...)
}

func (b Brush[color]) Repaint(p ...painted) painted {
	var result = b.Paint("")

	for i := range p {
		result.Append(p[i].content)
	}
	return result
}

func (b Brush[color]) Print(s ...string) {
	fmt.Print(b.Paint(s...))
}

func (b Brush[color]) Println(s ...string) {
	b.Print(strings.Join(s, " "), "\n")
}

func (b Brush[color]) Embed(values ...any) []painted {
	var (
		result  []painted
		current *painted
		add     = func(s string) {
			if s == "" {
				return
			}

			if current != nil {
				current.Append(s)
				return
			}

			result = append(result, b.Paint(s))
			current = &result[len(result)-1]
		}
	)

	for _, rawValue := range values {
		switch v := rawValue.(type) {
		case painted:
			result = append(result, v)
			current = nil
			continue
		case string:
			add(v)
		case fmt.Stringer:
			add(v.String())
		default:
			add(fmt.Sprint(v))
		}
	}

	return result
}
