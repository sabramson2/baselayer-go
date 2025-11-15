package baselayergo

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

//----------------------------------------
/*
create a string of random digits 0-9, of length n
*/
func RandNumString(n int) string {
	var sb strings.Builder
	sb.Grow(n)
	for range n {
		val := rand.IntN(10)
		fmt.Fprintf(&sb, "%d", val)
	}
	return sb.String()
}

//----------------------------------------
/*
get a random number between 0 (inclusive) and max (exclusive)
*/
func RandNum(max int) int {
	return rand.IntN(max)
}