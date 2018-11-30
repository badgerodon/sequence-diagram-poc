package parser

//go:generate pigeon -o parser.peg.go parser.peg

// A SendOp is a send operation
type SendOp struct {
	Src, Dst, Message string
}

// ParseDiagram parses a diagram
func ParseDiagram(src string) ([]SendOp, error) {
	ops, err := Parse("diagram", []byte(src))
	if err != nil {
		return nil, err
	}

	var sops []SendOp
	for _, o := range ops.([]interface{}) {
		sops = append(sops, o.(SendOp))
	}
	return sops, nil
}
