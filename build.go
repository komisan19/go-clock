package main

import (
	"fmt"
	"time"
)

func clock() {
	now := time.Now()
	fmt.Println(now)
}

func main() {
	clock()
}
