package brush

import (
	"fmt"
	"strings"
)

const (
	esc        = '\x1b'
	csi        = string(esc) + "["
	colorReset = "0"
)

// Painted represents a string that contains information about it's foreground and background color output
type Painted struct {
	foreground, background string
	content                string
}

// Paint some strings (joined without separator) with the specified font and background color
func Paint[color ColorType](font color, background Optional[color], s ...string) Painted {
	return Painted{
		foreground: font.foreground(),
		background: serializeBg(background),
		content:    strings.Join(s, ""),
	}
}

// String gives a string that contains some special sequence that will apply styling
func (p Painted) String() string {
	style := p.foreground
	if len(p.background) > 0 {
		style += ";" + p.background
	}

	return fmt.Sprintf("%s%sm%s%sm", csi, style, p.content, csi+colorReset)
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
