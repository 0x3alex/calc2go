package calc2go

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	k := tokenize("1 + 1")
	for _, v := range k {
		fmt.Printf("Token Type %d with value %s\n", v.tokenType, v.value)
	}
	tr, _ := covertToTreeNodes(k, 0)
	tr.printTree()
	fmt.Printf("%f\n", eval(tr))
}
