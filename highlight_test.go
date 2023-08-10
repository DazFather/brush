package brush_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

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

func ExampleHighlighted_Append() {
	var (
		marker = brush.New(brush.Black, brush.UseColor(brush.Yellow))
		h      = marker.Highlight("Hello world!", regexp.MustCompile("Hello"))
	)

	fmt.Print(h.Append(" ", marker.Paint("Hi"), " everyone!"))
	// Output: [30;43mHello[0m world! [30;43mHi[0m everyone!
}

func ExampleBrush_Embed() {
	var (
		green  = brush.New(brush.Green, nil)
		blue   = brush.New(brush.Blue, nil)
		yellow = brush.New(brush.Yellow, nil)
	)

	fmt.Print(blue.Embed(
		yellow.Highlight("Sun is yellow\n", regexp.MustCompile(`Sun|yellow`)),
		green.Paint("Grass"), " is ", green.Paint("green"),
		"\nAll the rest is blue",
	))
	// Output:
	// [33mSun[0m[34m is [0m[33myellow[0m[34m
	// [0m[32mGrass[0m[34m is [0m[32mgreen[0m[34m
	// All the rest is blue[0m
}

/* ---[ TESTS ]--- */

func TestBrush_Highlight(t *testing.T) {
	var (
		rgx    = regexp.MustCompile(`red`)
		marker = brush.New(brush.Red, nil)
	)

	assert(t, `Highlighting: "red"`,
		marker.Highlight("trashredgarbage", rgx).String(),
		"trash[31mred[0mgarbage",
	)

	assert(t, `Highlighting: "red banana"`,
		marker.Highlight("trash red banana trash red lolred", rgx).String(),
		"trash [31mred[0m banana trash [31mred[0m lol[31mred[0m",
	)
}

func TestHighlighted_Append(t *testing.T) {
	var (
		rgx    = regexp.MustCompile(`red( \w+)?`)
		marker = brush.New(brush.Red, nil)
	)

	banana := marker.Highlight("trash red banana", rgx)
	assert(t, `Append to "banana"`,
		banana.Append(" trash").String(),
		"trash [31mred banana[0m trash",
	)

	test := marker.Highlight("trashredgarbage", rgx)
	assert(t, `Append`,
		test.Append(
			cool{},
			"ciao",
			brush.Paint(brush.Green, nil, "green"),
			3,
			&banana,
		).String(),
		"trash[31mred[0mgarbagecoolciao[32mgreen[0m3trash [31mred banana[0m trash",
	)
}
