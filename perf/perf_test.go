package perf

import (
	f "fmt"
	"testing"
)

func someFunc1() {
	var x int = 2
	for i := range 10000000 {
		x *= x + i
	}
}



func TestTimeItSingle(t *testing.T) {
	r := TimeItSingle(func() {
		var x int = 2
		for i := range 10000000 {
			x *= x + i
		}
	})
	f.Printf("r = %d\n", r)
}

func TestTimeItMany(t *testing.T) {
	r := TimeItMany(20, someFunc1)
	r.PrintWithValues()
}

func TestTimeItMany2(t *testing.T) {
	r := TimeItMany(1001, someFunc1)
	r.Print()
}

