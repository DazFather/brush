package main

import (
	"fmt"

	"github.com/DazFather/brush"
)

func main() {
	var (
		mybrush     = brush.New(brush.BrightWhite.ToExtended(), nil)
		percentages = map[string]int{ // following percentages are not real
			"Cat":    65,
			"Dog":    45,
			"Parrot": 23,
			"Ants":   1,
		}
	)

	// Create and print the scale
	var scale = make([]brush.ExtendedANSIColor, brush.MaxIntensity+1)
	fmt.Print("scale: ")
	for i := brush.ZeroIntensity; i <= brush.MaxIntensity; i++ {
		color := brush.RGB(
			i,
			brush.MaxIntensity-i,
			brush.ZeroIntensity,
		)
		scale[i] = color
		fmt.Print(mybrush.UseBgColor(color).Paint(" ", i*20, "% "))
	}
	mybrush.UseDefaultColor()
	fmt.Print("\n\n")

	// Show result accordingly
	for name, prc := range percentages {
		intensity := len(scale) * prc / 100

		fmt.Println(
			name,
			"lovers:\t",
			mybrush.UseFontColor(scale[intensity]).Paint(prc, "%"),
		)
	}

}
