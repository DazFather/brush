package brush

import "fmt"

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

// Print shows on stdout some values (joined without separator)
// enforcing the current font and background color of the brush
func (b Brush[color]) Print(values ...any) {
	fmt.Print(b.Paint(values...))
}

// Println shows on stdout some values (joined with " ", and adding a "\n" at the end)
// enforcing the current font and background color of the brush
func (b Brush[color]) Println(values ...any) {
	fmt.Print(b.Paintln(values...))
}
