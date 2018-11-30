package main

import (
	"syscall/js"
)

func main() {
	svg := js.Global().Get("document").Call("createElementNS", "http://www.w3.org/2000/svg", "svg")
	svg.Call("setAttribute", "width", "600")
	svg.Call("setAttribute", "height", "200")
	js.Global().Get("document").Call("getElementsByTagName", "body").Index(0).Call("appendChild", svg)
}
