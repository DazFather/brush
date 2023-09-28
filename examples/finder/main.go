package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/DazFather/brush"
)

func main() {
	var err error

	if len(os.Args) == 3 {
		err = highlight(os.Args[1], os.Args[2])
	} else {
		err = fmt.Errorf("invalid given argument, expected: <pattern> <filename>")
	}

	if err != nil {
		fmt.Print("[", brush.Paint(brush.Red, nil, "ERROR"), "]: ", err)
	}
}

func highlight(pattern, filename string) (err error) {
	var (
		marker  = brush.New(brush.Black, brush.UseColor(brush.Yellow))
		content []byte
		rgx     *regexp.Regexp
	)

	if rgx, err = regexp.Compile(pattern); err == nil {
		if content, err = os.ReadFile(filename); err == nil {
			fmt.Println(marker.Highlight(string(content), rgx))
		}
	}

	return
}
