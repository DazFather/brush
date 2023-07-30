package brush

import (
	"fmt"
	"strings"
)

// Brush lets you paint some strings and change styling more freely
type Brush[color ColorType] struct {
	Foreground, defForeground color
	Background, defBackground Optional[color]
}

// New creates a new Brush with the given default colors of a specified set
func New[color ColorType](font color, background Optional[color]) Brush[color] {
	var b = Brush[color]{
		defForeground: font,
		defBackground: background,
	}
	b.UseDefaultColor()

	return b
}

// UseFontColor overrides the font color and gives back the same (now modified) brush
func (b *Brush[color]) UseFontColor(c color) *Brush[color] {
	b.Foreground = c
	return b
}

// UseFontColor overrides the background color and gives back the same (now modified) brush
func (b *Brush[color]) UseBgColor(c color) *Brush[color] {
	b.Background = UseColor(c)
	return b
}

// UseFontColor overrides the background color removing it and gives back the same (now modified) brush
func (b *Brush[color]) UseBgTransparent() *Brush[color] {
	b.Background = nil
	return b
}

// Swap overrides font and background color by inverting them and gives back the same (now modified) brush
// If background was unset then the default foreground color will be used as foreground
func (b *Brush[color]) Swap() *Brush[color] {
	b.Background, b.Foreground = UseColor(b.Foreground), PickColor(b.Background, b.defForeground)
	return b
}

// UseDefaultColor overrides font and background color by using the default values and gives back the same (now modified) brush
func (b *Brush[color]) UseDefaultColor() *Brush[color] {
	b.Foreground, b.Background = b.defForeground, b.defBackground
	return b
}

// Paint some strings (joined without separator) with the current font and background color of the brush
func (b Brush[color]) Paint(s ...string) Painted {
	return Paint(b.Foreground, b.Background, s...)
}

// Repaint some previously Painted items joining their contents (without separator)
// and using the current font and background color of the brush
func (b Brush[color]) Repaint(p ...Painted) Painted {
	var result = b.Paint("")

	for i := range p {
		result.Append(p[i].content)
	}
	return result
}

// Print shows on stdout some strings (joined without separator)
// applying the current font and background color of the brush
func (b Brush[color]) Print(s ...string) {
	fmt.Print(b.Paint(s...))
}

// Println shows on stdout some strings (joined with " ") and adding a "\n" at the end
// applying the current font and background color of the brush
func (b Brush[color]) Println(s ...string) {
	b.Print(strings.Join(s, " "), "\n")
}

// Embed lets create a list of Painted items by joining the given values into as few as possible
// applying the current font and background color of the brush
// If a value is a Painted item it gets added to the list of result and it's style it maintained
func (b Brush[color]) Embed(values ...any) []Painted {
	var (
		result  []Painted
		current *Painted
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
		case Painted:
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
