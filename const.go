package main

/*
Get the Max and the Min size of int from the system
*/
const (
	BitsPerWord = 32 << (^uint(0) >> 63)
	MaxInt  = 1<<(BitsPerWord-1) - 1
	MinInt  = -MaxInt - 1
)
