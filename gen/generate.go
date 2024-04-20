package gen

import (
	"fmt"
	"strings"

	"bfc/lex"
	"bfc/parse"
)

func Generate(ins []parse.Instruction) []byte {
	sb := strings.Builder{}
	sb.WriteString("_main:\n")
	for i, op := range ins {
		switch op.Op {
		case lex.Add:
			sb.WriteString("\tAdd\n")
		case lex.Sub:
			sb.WriteString("\tSub\n")
		case lex.Righ:
			sb.WriteString("\tRight\n")
		case lex.Left:
			sb.WriteString("\tLeft\n")
		case lex.JumpBack:
			sb.WriteString(fmt.Sprintf("\tb %s:\n", label(op.V)))
		case lex.JumpTo:
			sb.WriteString(fmt.Sprintf("%s:\n", label(i)))
		case lex.Print:
			sb.WriteString("\tPrint\n")
		case lex.Read:
			sb.WriteString("\tRead\n")

		}
	}

	return []byte(sb.String())
}

func label(i int) string {
	return fmt.Sprintf("i%d", i)
}
