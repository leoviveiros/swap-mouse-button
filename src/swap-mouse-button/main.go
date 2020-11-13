package main

import (
	"fmt"
	"syscall"
)

const (
	RESTORE = 0
	SWAP = 1
)

func main() {
	var mod = syscall.NewLazyDLL("user32.dll")
	var proc = mod.NewProc("SwapMouseButton")
	var fSwap = SWAP // TODO change via arg --restore

	ret, _, _ := proc.Call(uintptr(fSwap))

	fmt.Printf("Return: %d\n", ret)
}