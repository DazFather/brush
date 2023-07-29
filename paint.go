package brush

import (
	"fmt"

	"github.com/muesli/termenv"
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

	return fmt.Sprintf("%s%sm%s%sm", termenv.CSI, style, p.origin, termenv.CSI+termenv.ResetSeq)
}
