package main

import (
	"fmt"
	"github.com/DazFather/brush"
)

func main() {
	magenta := brush.New(brush.Magenta, nil)
	brush.Disable = true
	fmt.Println(magenta.Embed("dioporco"))
}
