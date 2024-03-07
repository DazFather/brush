package brush_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DazFather/brush"
)

/* ---[ EXAMPLES ]--- */

func ExamplePaint() {
	fmt.Println("I", brush.Paint(brush.Red, nil, "love"), "go")

	// Output:
	// I [31mlove[0m go
}

func ExamplePaintln() {
	painted := brush.Paintln(brush.Black, brush.UseColor(brush.White),
		"Hello", "World", "!",
	)
	fmt.Print(painted)

	// Output:
	// [30;47mHello World !
	//[0m
}

func ExamplePaintf() {
	blue := brush.Paintf(brush.White, brush.UseColor(brush.Blue), "%s", "blue")
	fmt.Println("The sky is", blue)

	// Output:
	// The sky is [37;44mblue[0m
}

func ExampleBrush_Paint() {
	myBrush := brush.New(brush.Red, nil)
	fmt.Println("I", myBrush.Paint("love"), "go")

	// Output:
	// I [31mlove[0m go
}

func ExampleBrush_Paintln() {
	myBrush := brush.New(brush.Black, brush.UseColor(brush.White))
	fmt.Print(myBrush.Paintln("Hello", "World", "!"))

	// Output:
	// [30;47mHello World !
	//[0m
}

func ExampleBrush_Paintf() {
	myBrush := brush.New(brush.White, brush.UseColor(brush.Blue))
	fmt.Println("The sky is", myBrush.Paintf("%s", "blue"))

	// Output:
	// The sky is [37;44mblue[0m
}

func ExamplePainted_Append() {
	banans := brush.Paint(brush.Yellow, nil, "banana")
	fmt.Println(banans.Append("s"))

	// Output: [33mbananas[0m
}

// Prepend a string at the start of the content of the painted item
// Warning: Do not use string containing styling
func ExamplePainted_Prepend() {
	banana := brush.Paint(brush.Yellow, nil, "banana")
	fmt.Println(banana.Prepend("yellow "))

	// Output: [33myellow banana[0m
}

// Replace the content of the painted item with another string
// Is possible to use the %s to refer to embed previous content
// Warning: Do not use string containing styling
func ExamplePainted_Replace() {
	banana := brush.Paint(brush.Yellow, nil, "banana")
	fmt.Println(banana.Replace(`The name "%s", is funny`))

	// Output: [33mThe name "banana", is funny[0m
}

/* ---[ TESTS ]--- */

func TestPaint(t *testing.T) {
	assert(t, `Paiting "I love go"`,
		fmt.Sprintln("I", brush.Paint(brush.Red, nil, "love"), "go"),
		"I [31mlove[0m go\n",
	)

	assert(t, `Paiting "banana"`,
		brush.Paint(brush.Yellow, nil, "banana").String(),
		"[33mbanana[0m",
	)

	ikea := brush.New(brush.Blue, brush.UseColor(brush.Yellow)).
		Highlight("IKEA (not really) logo", regexp.MustCompile("IKEA"))
	assert(t, `Paiting "ikea"`,
		brush.Paint(brush.Yellow, brush.UseColor(brush.Blue), 1, " ", cool{}, " reversed ", ikea).String(),
		"[33;44m1 cool reversed IKEA (not really) logo[0m",
	)
	assert(t, `Paiting "ikea" only`,
		brush.Paint(brush.White.ToTrueColor(), nil, &ikea).String(),
		"[38;2;192;192;192mIKEA (not really) logo[0m",
	)
}

func TestPaintln(t *testing.T) {
	banana := brush.Paint(brush.Green, nil, "banana")

	assert(t, `Paiting "3 cool yellow bananas"`,
		brush.Paintln(brush.Black, brush.UseColor(brush.Yellow), 3, cool{}, "yellow", banana.Append("s")).String(),
		"[30;43m3 cool yellow bananas\n[0m",
	)

	assert(t, `Painting ""`,
		brush.Paintln(brush.White, nil, "").String(),
		"[37m\n[0m",
	)
}

/* ---[ UTILS ]--- */

func assert[T comparable](t *testing.T, prefix string, got, want T) (pass bool) {
	if pass = want == got; !pass {
		t.Error(prefix, "| want: ", want, " got: ", got)
	}

	return
}

type cool struct{}

func (c cool) String() string {
	return "cool"
}
