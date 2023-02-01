package main

import "txp/web-service-gin/app"

func main() {
	app := &app.App{}
	app.Init()
	app.Run()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
