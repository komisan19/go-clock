package main

import (
	"syscall/js"
	"time"
)

func clock() {
	now := time.Now()

	date := js.Global().Get("Date").New(now.Unix() * 1000)
	js.Global().Get("document").Call("getElementById", "clock").Set("innerHTML", date)
}

func main() {
	clock()
	select {}
}
