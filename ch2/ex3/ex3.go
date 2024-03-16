package main

import (
	"fmt"
	"math"
)

func main() {
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI int64 = math.MaxInt64
	fmt.Printf("b: %d\nsmallI: %d\nbigI: %d\n", b, smallI, bigI)
}
