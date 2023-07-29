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

type painted struct {
	foreground, background string
	content                string
}

func Paint[color ColorType](font color, background Optional[color], s ...string) painted {
	return painted{
		foreground: font.foreground(),
		background: serializeBg(background),
		content:    strings.Join(s, ""),
	}
}

func (p painted) String() string {
	style := p.foreground
	if len(p.background) > 0 {
		style += ";" + p.background
	}

	return fmt.Sprintf("%s%sm%s%sm", csi, style, p.content, csi+colorReset)
}

func (p *painted) Append(s string) *painted {
	p.content += s
	return p
}

func (p *painted) Prepend(s string) *painted {
	p.content = s + p.content
	return p
}

func (p *painted) Replace(s string) *painted {
	p.content = strings.ReplaceAll(s, "%s", p.content)
	return p
}
