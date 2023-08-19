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
	marker := brush.New(brush.Black, brush.UseColor(brush.BrightYellow))

	fmt.Print(marker.HighlightFunc(
		"this is uppercase and yellow",
		regexp.MustCompile("uppercase"),
		strings.ToUpper,
	))
	// Output: this is [30;103mUPPERCASE[0m and yellow
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

	assert(t, `Highlighting: nothing`,
		marker.Highlight("garbage", rgx).String(),
		"garbage",
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

func TestBrush_HighlightFunc(t *testing.T) {
	var (
		rgx    = regexp.MustCompile(`red`)
		marker = brush.New(brush.Red, nil)
		repl   = func(_ string) string { return "" }
		text   = "garbage"
	)

	assert(t, `Highlighting: nothing`,
		marker.HighlightFunc(text, rgx, repl).String(),
		text,
	)
}

func TestHighlighted_Append(t *testing.T) {
	var (
		rgx    = regexp.MustCompile(`red( \w+)?`)
		marker = brush.New(brush.Red, nil)
	)

	banana := marker.Highlight("trash red banana", rgx)
	expected := "trash [31mred banana[0m"
	assert(t, `Append: nothing`,
		banana.Append().String(),
		expected,
	)

	expected += " trash"
	assert(t, `Append to "banana"`,
		banana.Append(" trash").String(),
		"trash [31mred banana[0m trash",
	)

	ciao := marker.Highlight("ciao", regexp.MustCompile(`(?i)[aeiou]`))

	test := marker.Highlight("trashredgarbage", rgx)
	assert(t, `Append`,
		test.Append(
			cool{},
			ciao,
			brush.Paint(brush.Green, nil, "green"),
			3,
			&banana,
		).String(),
		"trash[31mred[0mgarbagecoolc[31mi[0m[31ma[0m[31mo[0m[32mgreen[0m3trash [31mred banana[0m trash",
	)
}

func TestBrush_Embed(t *testing.T) {
	var (
		myBrush = brush.New(brush.Red, nil)
		marker  = brush.New(brush.Black, brush.UseColor(brush.Yellow))
	)

	assert(t, `Embedding: nothing`, myBrush.Embed().String(), "")

	banana := marker.Paint("banana")
	assert(t, `Embedding: "1 cool banana"`,
		myBrush.Embed(1, " ", cool{}, " ", banana).String(),
		"[31m1[0m[31m [0m[31mcool[0m[31m [0m[30;43mbanana[0m",
	)

	isYellow := brush.Join(" is ", marker.Paint("yellow"))
	assert(t, `Embedding: "The banana is yellow"`,
		myBrush.Embed("The ", &banana, isYellow).String(),
		"[31mThe [0m[31m[30;43mbanana[0m[0m[31m is [0m[30;43myellow[0m",
	)

	text := "A fox jumps over the lazy dog"
	vouels := regexp.MustCompile(`(?i)[aeiou]`)
	h := marker.Highlight(text, vouels)
	expected := "[30;43mA[0m[31m f[0m[30;43mo[0m[31mx j[0m[30;43mu[0m[31mmps [0m[30;43mo[0m[31mv[0m[30;43me[0m[31mr th[0m[30;43me[0m[31m l[0m[30;43ma[0m[31mzy d[0m[30;43mo[0m[31mg[0m[31m![0m"

	assert(t, fmt.Sprintf(`Embedding: "%s"`, text),
		myBrush.Embed(h, "!").String(),
		expected,
	)

	assert(t, fmt.Sprintf(`Embedding: pointer to "%s"`, text),
		myBrush.Embed(&h, "!").String(),
		expected,
	)

	text = reverse(text)
	h = marker.Highlight(text, vouels)
	expected = "[31mg[0m[30;43mo[0m[31md yz[0m[30;43ma[0m[31ml [0m[30;43me[0m[31mht r[0m[30;43me[0m[31mv[0m[30;43mo[0m[31m spm[0m[30;43mu[0m[31mj x[0m[30;43mo[0m[31mf [0m[30;43mA[0m[31m![0m"
	assert(t, fmt.Sprintf(`Embedding: "%s"`, text),
		myBrush.Embed(h, "!").String(),
		expected,
	)

	assert(t, fmt.Sprintf(`Embedding: pointer to "%s"`, text),
		myBrush.Embed(&h, "!").String(),
		expected,
	)
}

func reverse(s string) string {
	size := len(s)
	runes := make([]rune, size)
	for i, ch := range s {
		runes[size-1-i] = ch
	}
	return string(runes)
}
