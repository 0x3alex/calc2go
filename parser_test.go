package calc2go

import (
	"testing"
)

func TestParser(t *testing.T) {
	println(Eval("1 + (5*6.2+(10^(3*10.5)))"))
}
