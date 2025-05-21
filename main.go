package main

import (
	"fmt"
)

// function to be called by Lua
func adder(a1 float64, a2 float64) float64 {
	return a1 + a2
}

func main() {
	ctx, err := NewContext()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if err := ctx.LoadFile("b.lua", map[string]interface{}{
		"adder": adder,
	}); err == nil {

	} else {
		fmt.Println(err)
	}
}
