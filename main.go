package main

import "txp/web-service-gin/src"

func main() {
	app := &src.App{}
	app.Init()
	app.Run()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
