package baselayergo

import (
	"testing"
)

func TestCodeCreate(t *testing.T) {
	//t.Logf("file: %s\n", CodeCreate())
	s, e := CodeCreate(100)
	if e != nil { t.Fail(); return }
	Pf("code = '%s'\n", s)
	r, e := CodeVerify(s)
	if e != nil { t.Fail(); return }
	Pf("result = %t\n", r)
}

func TestRand(t *testing.T) {
	for range(10) {
		Pf("r = %s\n", RandNumString(5))
	}
}