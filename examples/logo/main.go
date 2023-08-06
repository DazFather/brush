package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/DazFather/brush"
)

func main() {
	var (
		content = brush.New(brush.BrightBlack, brush.UseColor(brush.Black))
		rgx     = regexp.MustCompile(`[A-Z\.]`)
		title   = [...]string{
			" brush brush brush brush brush brush brush brush ",
			"bruSH.BRush.BRUSh brUSh bRUsh BRUSh BRush.Brush",
			"rusH.brUSh BRusH.bruSH brUSh BRush bRUsh BRush ",
			"ush.BRUSh bRUSH.brusH.bruSH brUSH brUSH.BRUsh brush",
			"sh BRusH.brUSh BRush.BrusH.brush.BruSH brUSh brush",
			"h bRUSH.bruSH brUSh bRUSH.bruSH.BrusH.bruSH ",
			" brush brush brush brush brush brush brush brush",
		}
	)

	for i, line := range title {
		color := brush.ANSIColor(i + 1)
		rainbow := brush.New(color+8, &color)

		fmt.Println(rainbow.Embed(content.HighlightFunc(line, rgx, dotToSpace)))
	}
}

func dotToSpace(s string) string {
	return strings.ReplaceAll(s, ".", " ")
}
