package calc2go

import (
	"math"
	"strconv"
)

func isLeaf(node *ASTNode) bool {
	return node.l == nil && node.r == nil
}

func eval(root *ASTNode) float64 {
	if isLeaf(root) {
		v, err := strconv.ParseFloat(root.t.value, 32)
		if err != nil {
			v, err := strconv.Atoi(root.t.value)
			if err != nil {
				panic("Error while eval with value " + root.t.value)
			}
			return float64(v)
		}
		return v
	}
	left := eval(root.l)
	right := eval(root.r)
	switch root.t.tokenType {
	case PLUS:
		return left + right
	case MINUS:
		return left - right
	case MULTIPLY:
		return left * right
	case DIVIDE:
		return left / right
	case EXP:
		return math.Pow(left, right)
	}
	panic("UNREACHABLE " + root.t.value)
	return 0
}
