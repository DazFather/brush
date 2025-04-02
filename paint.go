package brush

import (
	"fmt"
	"strings"

	"golang.org/x/term"
	"os"
)

// Painted represents a string that contains information about it's foreground and background color output
type Painted struct {
	content string
	disable bool
	style
}

var (
	isATTY = term.IsTerminal(int(os.Stdout.Fd()))

	// DisableIfNotTTY disables Painted and Brush if it detects application is
	// not in a tty at a global level (for this library)
	DisableIfNotTTY = true

	// Disable colored output. By default it depends by DisableIfNotTTY
	Disable = DisableIfNotTTY && !isATTY
)

// Paint some values (joined without separator) with the specified font and background color.
// If a Painted and/or an Highlighted item is given, they will lose their previous style
// and provided font and background colors will be enforced
func Paint[color ColorType](font color, background Optional[color], values ...any) Painted {
	var res = Painted{
		style:   serialize(font, background),
		disable: Disable || DisableIfNotTTY && !isATTY,
	}

	for _, v := range values {
		res.content += extractContent(v)
	}

	return res
}

// Paintln is like Paint but similarly to fmt.Sprintln it separates values with " "
// and it adds a "\n" at the end of all.
// If a Painted and/or an Highlighted item is given, they will lose their previous style
// and provided font and background colors will be enforced
func Paintln[color ColorType](font color, background Optional[color], values ...any) Painted {
	const separator = " "
	var res = Painted{
		style:   serialize(font, background),
		disable: Disable || DisableIfNotTTY && !isATTY,
	}

	if len(values) == 0 {
		res.content = "\n"
		return res
	}

	last := len(values) - 1
	for i := range values[:last] {
		res.content += extractContent(values[i]) + " "
	}
	res.content += extractContent(values[last]) + "\n"

	return res
}

// Paintf is like Paint but similarly to fmt.Sprintf it allows to use a model
// with some placeholders that will be replaced by the values.
// If a Painted and/or an Highlighted item is given, they will lose their previous style
// and provided font and background colors will be enforced
func Paintf[color ColorType](font color, background Optional[color], model string, values ...any) Painted {
	for i, v := range values {
		values[i] = extractContent(v)
	}

	return Painted{
		style:   serialize(font, background),
		content: fmt.Sprintf(model, values...),
	}
}

// Paint some values (joined without separator) with the current brush colors.
// If a Painted and/or an Highlighted item is given, they will lose their previous style
// and the current styling of the brush will be enforced
func (b Brush[color]) Paint(values ...any) Painted {
	p := Paint(b.Foreground, b.Background, values...)
	p.disable = b.Disable
	return p
}

// Paintln like Paint but similarly to fmt.Sprintln it separates values with " "
// and it adds a "\n" at the end of all.
// If a Painted and/or an Highlighted item is given, they will lose their previous style
// and the current styling of the brush will be enforced
func (b Brush[color]) Paintln(values ...any) Painted {
	p := Paintln(b.Foreground, b.Background, values...)
	p.disable = b.Disable
	return p
}

// Paintf is like Paint but similarly to fmt.Sprintf it allows to use a model
// with some placeholders that will be replaced by the values.
// If a Painted and/or an Highlighted item is given, they will lose their previous style
// and the current styling of the brush will be enforced
func (b Brush[color]) Paintf(model string, values ...any) Painted {
	p := Paintf(b.Foreground, b.Background, model, values...)
	p.disable = b.Disable
	return p
}

// String gives a string that contains some special sequence that will apply styling
func (p Painted) String() string {
	if p.disable {
		return p.content
	}
	return p.apply(p.content)
}

// Append a string at the end of the content of the painted item
// Warning: Do not use string containing styling
func (p *Painted) Append(s string) *Painted {
	p.content += s
	return p
}

// Prepend a string at the start of the content of the painted item
// Warning: Do not use string containing styling
func (p *Painted) Prepend(s string) *Painted {
	p.content = s + p.content
	return p
}

// Replace the content of the painted item with another string
// Is possible to use the %s to refer to embed previous content
// Warning: Do not use string containing styling
func (p *Painted) Replace(s string) *Painted {
	p.content = strings.ReplaceAll(s, "%s", p.content)
	return p
}

func extractContent(value any) (content string) {
	switch v := value.(type) {
	case Painted:
		content = v.content
	case *Painted:
		content = v.content
	case Highlighted:
		content = v.content
	case *Highlighted:
		content = v.content
	case string:
		content = v
	case fmt.Stringer:
		content = v.String()
	default:
		content = fmt.Sprint(v)
	}

	return
}
