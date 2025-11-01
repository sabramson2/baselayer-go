package perf

import (
	f "fmt"
	"testing"

	bl "github.com/sabramson2/baselayer-go"
)

func someFunc1() {
	var x int = 2
	for i := range 10000000 {
		x *= x + i
	}
}



func TestTimeItSingle(t *testing.T) {
	r := bl.TimeItSingle(func() {
		var x int = 2
		for i := range 10000000 {
			x *= x + i
		}
	})
	f.Printf("r = %d\n", r)
}

func TestTimeItMany(t *testing.T) {
	r := bl.TimeItMany(20, someFunc1)
	r.PrintWithValues()
}

func TestTimeItMany2(t *testing.T) {
	r := bl.TimeItMany(1001, someFunc1)
	r.Print()
}

