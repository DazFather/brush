package brush

import "fmt"

const (
	esc        = '\x1b'
	csi        = string(esc) + "["
	colorReset = "0"
)

type painted struct {
	foreground, background string
	origin                 string
}

func (p painted) String() string {
	style := p.foreground
	if len(p.background) > 0 {
		style += ";" + p.background
	}

	return fmt.Sprintf("%s%sm%s%sm", csi, style, p.origin, csi+colorReset)
}
