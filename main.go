package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	bitmap := robotgo.CaptureScreen(100, 100, 100, 100)
	x, y := robotgo.FindBitmap(bitmap)
	fmt.Println("...", x, y)
}
