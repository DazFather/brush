package brush

import "strings"

// Painted represents a string that contains information about it's foreground and background color output
type Painted struct {
	content string
	style
}

// Paint some strings (joined without separator) with the specified font and background color
func Paint[color ColorType](font color, background Optional[color], s ...string) Painted {
	return Painted{
		content: strings.Join(s, ""),
		style:   serialize(font, background),
	}
}

// Paint some strings (joined without separator) with the current font and background color of the brush
func (b Brush[color]) Paint(s ...string) Painted {
	return Painted{
		content: strings.Join(s, ""),
		style:   b.extract(),
	}
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

// String gives a string that contains some special sequence that will apply styling
func (p Painted) String() string {
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
