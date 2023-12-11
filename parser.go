package calc2go

import (
	"strings"
	"unicode"
)

const (
	UNKNOWN = -1
	PLUS    = iota
	MINUS
	MULTIPLY
	DIVIDE
	NUM
	OPEN
	CLOSE
	EXP
)

var typeValMap = map[rune]int{
	'+': PLUS,
	'-': MINUS,
	'*': MULTIPLY,
	'/': DIVIDE,
	'(': OPEN,
	')': CLOSE,
	'^': EXP,
}

type token struct {
	value     string
	tokenType int
}

type ASTNode struct {
	t token
	l *ASTNode
	r *ASTNode
}

func (a *ASTNode) printTree() {
	if a == nil {
		return
	}
	a.l.printTree()
	println("Value ", a.t.value)
	a.r.printTree()
}

func (a *ASTNode) size() int {
	if a == nil {
		return 0
	}
	return a.l.size() + 1 + a.r.size()
}

func getTokenType(a rune) int {
	if unicode.IsDigit(a) {
		return NUM
	}
	if v, ok := typeValMap[a]; ok {
		return v
	}
	return UNKNOWN
}

func covertToTreeNodes(tokens []token, i int) (*ASTNode, int) {
	//var nodes []ASTNode
	println("started with ", i)
	var nodes []ASTNode
	mergeNodes := func() ASTNode {
		return ASTNode{
			t: token{
				value:     nodes[1].t.value,
				tokenType: nodes[1].t.tokenType,
			},
			l: &ASTNode{
				t: nodes[0].t,
				l: nodes[0].l,
				r: nodes[0].r,
			},
			r: &ASTNode{
				t: nodes[2].t,
				l: nodes[2].l,
				r: nodes[2].r,
			},
		}
	}
	for ; i < len(tokens); i++ {
		if tokens[i].tokenType == CLOSE {
			println("closing with i=", i)
			break
		}
		if tokens[i].tokenType == OPEN {
			n, t := covertToTreeNodes(tokens, i+1)
			i = t
			//println("Now at i ", i, " with length ", len(tokens))
			nodes = append(nodes, *n)
		} else {
			nodes = append(nodes, ASTNode{t: tokens[i]})
		}
		if len(nodes) == 3 {
			nn := mergeNodes()
			nodes = nil
			nodes = append(nodes, nn)
		}
	}
	//println("len of nodes after loop is ", len(nodes))
	if len(nodes) == 3 {
		nn := mergeNodes()
		nodes = nil
		nodes = append(nodes, nn)
	}
	return &nodes[0], i
}

func tokenize(s string) []token {
	s = strings.ReplaceAll(s, " ", "") //we don't want them
	var tokens []token
	currentToken := token{}
	saveAndCreateToken := func(v rune) {
		tokens = append(tokens, currentToken)
		currentToken = token{}
		currentToken.value += string(v)
		currentToken.tokenType = getTokenType(v)
	}
	for _, v := range s {
		if len(currentToken.value) == 0 {
			currentToken.value += string(v)
			currentToken.tokenType = getTokenType(v)
			continue
		}
		/*
			To avoid )) or (( being one token
		*/
		if (currentToken.tokenType == OPEN && v == '(') || (currentToken.tokenType == CLOSE && v == ')') {
			saveAndCreateToken(v)
			continue
		}
		/*
			To allow float numbers
		*/
		if (currentToken.tokenType == NUM && v == '.') || currentToken.tokenType == getTokenType(v) {
			currentToken.value += string(v)
			continue
		}
		if getTokenType(v) != currentToken.tokenType {
			saveAndCreateToken(v)
		}
	}
	tokens = append(tokens, currentToken)
	return tokens
}
