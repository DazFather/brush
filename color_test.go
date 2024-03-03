package brush_test

import (
	"fmt"

	"github.com/DazFather/brush"
)

/* ---[ EXAMPLES ]--- */

func ExampleRGB() {
	pink := brush.RGB(
		brush.MaxIntensity,
		brush.MediumIntensity,
		brush.HightIntensity,
	)

	fmt.Println(brush.Paint(0, &pink, "Flamingo"))
	// Output: [38;5;0;48;5;218mFlamingo[0m
}

func ExampleGrayScale() {
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
	selectedBg := brush.UseColor(brush.Magenta)

	fmt.Println(brush.Paint(brush.BrightMagenta, selectedBg, "Magenta"), "is cool")
	// Output: [95;45mMagenta[0m is cool
}

func ExamplePickColor() {
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
	var (
		pinkish  = brush.TrueColor{Red: 255, Green: 82, Blue: 197}
		brownish = brush.TrueColor{155, 106, 0}
		myBrush  = brush.New(pinkish, &brownish)
	)

	myBrush.Println("HELLO WORLD")
	// Output: [38;2;255;82;197;48;2;155;106;0mHELLO WORLD
	// [0m
}
