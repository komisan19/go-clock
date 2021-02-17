package main

import (
	"syscall/js"
)

func clock(this js.Value, args []js.Value) interface{} {
	date := js.Global().Get("Date").New().Call("toLocaleString").String()

	js.Global().Get("document").Call("getElementById", "clock").Set("textContent", date)
	return nil
}

func registerCallbacks() {
	js.Global().Call("setInterval", js.FuncOf(clock), "200")
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
