package brush_test

import "github.com/DazFather/brush"

func ExampleBrush_UseFontColor() {
	brush.DisableIfNotTTY = false
	myBrush := brush.New(brush.Black, brush.UseColor(brush.White))

	myBrush.Println("something")
	myBrush.UseFontColor(brush.Blue).Println("something else")
	// Output:
	// [30;47msomething
	// [0m[34;47msomething else
	// [0m
}

func ExampleBrush_UseBgColor() {
	brush.DisableIfNotTTY = false
	myBrush := brush.New(brush.Black, brush.UseColor(brush.White))

	myBrush.Println("something")
	myBrush.UseBgColor(brush.Cyan).Println("something else")
	// Output:
	// [30;47msomething
	// [0m[30;46msomething else
	// [0m
}

func ExampleBrush_UseBgTransparent() {
	brush.DisableIfNotTTY = false
	myBrush := brush.New(brush.Black, brush.UseColor(brush.White))

	myBrush.Println("something")
	myBrush.UseBgTransparent().Println("something else")
	// Output:
	// [30;47msomething
	// [0m[30msomething else
	// [0m
}

func ExampleBrush_Swap() {
	brush.DisableIfNotTTY = false
	myBrush := brush.New(brush.Black, brush.UseColor(brush.White))

	myBrush.Println("something")
	myBrush.Swap().Println("something else")
	// Output:
	// [30;47msomething
	// [0m[37;40msomething else
	// [0m
}

func ExampleBrush_Print() {
	brush.DisableIfNotTTY = false
	myBrush := brush.New(brush.Black, brush.UseColor(brush.White))

	myBrush.Print("Hello", " World")
	// Output: [30;47mHello World[0m
}

func ExampleBrush_Println() {
	brush.DisableIfNotTTY = false
	myBrush := brush.New(brush.Black, brush.UseColor(brush.White))

	myBrush.Println("Hello", "World")
	// Output:
	// [30;47mHello World
	// [0m
}

func ExampleBrush_Printf() {
	brush.DisableIfNotTTY = false
	myBrush := brush.New(brush.Black, brush.UseColor(brush.White))

	myBrush.Printf("%s %s", "Hello", "World")
	// Output: [30;47mHello World[0m
}
