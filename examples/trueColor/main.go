package main

import (
	"fmt"

	"github.com/DazFather/brush"
)

func main() {
	var (
		pinkish = brush.TrueColor{255, 82, 197}
		brownish = brush.TrueColor{155, 106, 0}
		test = brush.New(pinkish, &brownish)
	)

	test.Println("Can you see this (correctly) ?")
	fmt.Println("If not, probably your terminal does not support true colors")
}
