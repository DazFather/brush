package brush_test

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/DazFather/brush"
)

/* ---[ EXAMPLES ]--- */

func ExampleJoin() {
	fmt.Print(brush.Join(
		brush.Paint(brush.Red, nil, "Roses are red"),
		",\n",
		brush.Paint(brush.Blue, nil, "Violets are blue"),
		",\nSugar is sweet,\nAnd so are you.",
	))
	// Output:
	// [31mRoses are red[0m,
	// [34mViolets are blue[0m,
	// Sugar is sweet,
	// And so are you.
}

func ExampleBrush_Highlight() {
	var (
		marker = brush.New(brush.Black, brush.UseColor(brush.Yellow))
		text   = "bla bla something intresting, something intresting blah bla"
		res    = marker.Highlight(text, regexp.MustCompile("something intresting"))
	)

	fmt.Print(res)
	// Output:
	// bla bla [30;43msomething intresting[0m, [30;43msomething intresting[0m blah bla
}

func ExampleBrush_HighlightFunc() {
	marker := brush.New(brush.Yellow, nil)

	fmt.Print(marker.HighlightFunc(
		"this is uppercase and yellow",
		regexp.MustCompile("uppercase"),
		strings.ToUpper,
	))
	// Output: this is [33mUPPERCASE[0m and yellow
}
