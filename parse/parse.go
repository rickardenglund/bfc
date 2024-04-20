package parse

import (
	"fmt"

	"bfc/lex"
)

type Instruction struct {
	Op lex.Op
	V  int
}

func (i Instruction) String() string {
	switch i.Op {
	case lex.Add:
		return "+"
	case lex.Sub:
		return "-"
	case lex.Righ:
		return ">"
	case lex.Left:
		return "<"
	case lex.JumpBack:
		return "["
	case lex.JumpTo:
		return fmt.Sprintf("(]%d)", i.V)
	case lex.Print:
		return "."
	case lex.Read:
		return ","
	default:
		return "?"

	}

}

func Parse(ts []lex.Token) ([]Instruction, error) {
	prog := []Instruction{}
	targets := []int{}
	for i, t := range ts {
		ins := Instruction{Op: t.Op}

		switch t.Op {
		case lex.JumpTo:
			targets = append(targets, i)

		case lex.JumpBack:
			if len(targets) == 0 {
				return nil, fmt.Errorf("missing target %d:%d", t.Row, t.Col)
			}

			ins.V = targets[len(targets)-1]
			targets = targets[:len(targets)-1]
		}

		prog = append(prog, ins)

	}

	if len(targets) > 0 {
		t := ts[targets[0]]
		return nil, fmt.Errorf("to many targets: %d:%d", t.Row, t.Col)
	}

	return prog, nil
}
