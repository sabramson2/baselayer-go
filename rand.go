package baselayergo

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func RandNumString(n int) string {
	var sb strings.Builder
	sb.Grow(n)
	for range n {
		val := rand.IntN(10)
		fmt.Fprintf(&sb, "%d", val)
	}
	return sb.String()
}