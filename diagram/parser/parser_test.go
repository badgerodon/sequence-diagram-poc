package parser

import "testing"

func Test(t *testing.T) {
	t.Error(ParseDiagram(`
	A -> B: test
`))
}
