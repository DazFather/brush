package brush_test

import (
	"fmt"
	"testing"

	"github.com/DazFather/brush"
)

/* ---[ EXAMPLES ]--- */

func ExampleRGB() {
	brush.DisableIfNotTTY = false // probably you don't want to override this variable

	pink := brush.RGB(
		brush.MaxIntensity,
		brush.MediumIntensity,
		brush.HightIntensity,
	)

	fmt.Println(brush.Paint(0, &pink, "Flamingo"))
	// Output: [38;5;0;48;5;218mFlamingo[0m
}

func ExampleGrayScale() {
	brush.DisableIfNotTTY = false // probably you don't want to override this variable

	var (
		white     = brush.GrayScale(brush.MaxGrayScale)
		lightGray = brush.GrayScale(brush.MaxGrayScale - 5)
		gray      = brush.GrayScale(brush.MaxGrayScale / 2)
		black     = brush.GrayScale(0)

		myBrush = brush.New(brush.Black.ToExtended(), nil)
	)

	myBrush.UseBgColor(white).Println("Sunny")
	myBrush.UseBgColor(lightGray).Println("Cloudy")
	myBrush.UseBgColor(gray).Println("Rainy")
	myBrush.UseBgColor(black).Println("WTF")

	// Output:
	// [38;5;0;48;5;15mSunny
	// [0m[38;5;0;48;5;251mCloudy
	// [0m[38;5;0;48;5;243mRainy
	// [0m[38;5;0;48;5;0mWTF
	// [0m
}

func ExampleUseColor() {
	brush.DisableIfNotTTY = false // probably you don't want to override this variable

	selectedBg := brush.UseColor(brush.Magenta)

	fmt.Println(brush.Paint(brush.BrightMagenta, selectedBg, "Magenta"), "is cool")
	// Output: [95;45mMagenta[0m is cool
}

func ExamplePickColor() {
	brush.DisableIfNotTTY = false // probably you don't want to override this variable

	myBrush := brush.New(brush.Yellow, nil)
	reversed := brush.New(
		brush.PickColor(myBrush.Background, brush.Black),
		brush.UseColor(myBrush.Foreground),
	)

	myBrush.Println("default brush example")
	reversed.Println("reversed brush example")
	// Output: [33mdefault brush example
	// [0m[30;43mreversed brush example
	// [0m
}

func ExampleTrueColor() {
	brush.DisableIfNotTTY = false // probably you don't want to override this variable

	var (
		pinkish  = brush.TrueColor{Red: 255, Green: 82, Blue: 197}
		brownish = brush.TrueColor{155, 106, 0}
		myBrush  = brush.New(pinkish, &brownish)
	)

	myBrush.Println("HELLO WORLD")
	// Output: [38;2;255;82;197;48;2;155;106;0mHELLO WORLD
	// [0m
}

func ExampleParseHex() {
	brush.DisableIfNotTTY = false // probably you don't want to override this variable

	var (
		color *brush.TrueColor
		err   error
	)

	// Example with full RGB format (# is optional)
	if color, err = brush.ParseHex("#FFA500"); err == nil {
		fmt.Println("Yellow:", *color)
	}

	// Example with shorthand RGB format (# is optional)
	if color, err = brush.ParseHex("F00"); err == nil {
		fmt.Println("Red:", *color)
	}

	// Bad examples
	_, err = brush.ParseHex("blue")
	fmt.Println(err)
	_, err = brush.ParseHex("daz")
	fmt.Println(err)

	// Output:
	// Yellow: {255 165 0}
	// Red: {255 0 0}
	// Cannot parse blue color: Invalid hex length, must be 3 or 6 digits long (excluding optional prefix '#')
	// Cannot parse daz color: strconv.ParseUint: parsing "z": invalid syntax
}

/* ---[ TESTS ]--- */

func TestANSIColor_ToTrueColor(t *testing.T) {
	brush.DisableIfNotTTY = false

	// Test cases
	tests := []struct {
		input    brush.ANSIColor
		expected brush.TrueColor
	}{
		{brush.Black, brush.TrueColor{0, 0, 0}},
		{brush.Red, brush.TrueColor{128, 0, 0}},
		{brush.Green, brush.TrueColor{0, 128, 0}},
		{brush.Yellow, brush.TrueColor{128, 128, 0}},
		{brush.Blue, brush.TrueColor{0, 0, 128}},
		{brush.Magenta, brush.TrueColor{128, 0, 128}},
		{brush.Cyan, brush.TrueColor{0, 128, 128}},
		{brush.White, brush.TrueColor{192, 192, 192}},
		{brush.BrightBlack, brush.TrueColor{128, 128, 128}},
		{brush.BrightRed, brush.TrueColor{255, 0, 0}},
		{brush.BrightGreen, brush.TrueColor{0, 255, 0}},
		{brush.BrightYellow, brush.TrueColor{255, 255, 0}},
		{brush.BrightBlue, brush.TrueColor{0, 0, 255}},
		{brush.BrightMagenta, brush.TrueColor{255, 0, 255}},
		{brush.BrightCyan, brush.TrueColor{0, 255, 255}},
		{brush.BrightWhite, brush.TrueColor{255, 255, 255}},
	}

	// Iterate through test cases
	for _, test := range tests {
		result := test.input.ToTrueColor()
		if result != test.expected {
			t.Errorf("ToTrueColor(%v): got %v, want %v", test.input, result, test.expected)
		}
	}
}

func TestExtendedANSIColor_ToTrueColor(t *testing.T) {
	brush.DisableIfNotTTY = false

	// Test cases
	tests := []struct {
		input    brush.ExtendedANSIColor
		expected brush.TrueColor
	}{
		{brush.Magenta.ToExtended(), brush.TrueColor{128, 0, 128}},
		{brush.RGB(brush.MaxIntensity, 0, 0), brush.BrightRed.ToTrueColor()},
		{brush.GrayScale(brush.MaxGrayScale), brush.BrightWhite.ToTrueColor()},
		{brush.GrayScale(21), brush.TrueColor{208, 208, 208}},
	}

	// Iterate through test cases
	for _, test := range tests {
		result := test.input.ToTrueColor()
		if result != test.expected {
			t.Errorf("ToTrueColor(%v): got %v, want %v", test.input, result, test.expected)
		}
	}
}
