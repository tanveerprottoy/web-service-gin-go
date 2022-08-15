package main

import "txp/web-service-gin/src"

func main() {
	application := &app.App{}
	application.Init()
	application.Run()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}