package gen

import (
	"fmt"
	"strings"

	"bfc/lex"
	"bfc/parse"
)

func Generate(ins []parse.Instruction) []byte {
	sb := strings.Builder{}
	sb.WriteString(head)
	for i, op := range ins {
		switch op.Op {
		case lex.Add:
			sb.WriteString(`//Add
	ldr x0, [x9, x10]
	add x0, x0, #1
	str x0, [x9, x10]
`)
		case lex.Sub:
			sb.WriteString(`//Sub
	ldr x0, [x9, x10]
	sub x0, x0, #1
	str x0, [x9, x10]
`)
		case lex.Righ:
			sb.WriteString(`//Right
	add x10, x10, #4
`)
		case lex.Left:
			sb.WriteString(`//Left
		sub x10, x10, #4`)
		case lex.JumpBack:
			sb.WriteString(fmt.Sprintf(`//Jump back
	ldr x0, [x9, x10]
	cmp x0, #0
	bne %s
`, label(op.V)))
		case lex.JumpTo:
			sb.WriteString(fmt.Sprintf("%s:\n", label(i)))
		case lex.Print:
			sb.WriteString(`//Print
	bl _print
`)
		case lex.Read:
			sb.WriteString("\tRead\n")
		}
	}

	sb.WriteString(foot)

	return []byte(sb.String())
}

func label(i int) string {
	return fmt.Sprintf("i%d", i)
}

const head = `
.global _start
.align 2

_start:
	adrp x9,  _mem@PAGE
	add x9, x9, _mem@PAGEOFF
	ldr x0, [x9]
	mov x10, #0
`

const foot = `
	//bl _print
_exit:
	mov X0, #0
	mov X16, #1
	svc 0
_print:
	mov X0, #1
	add x1, x9, x10
	mov X2, #1
	mov X16, #4
	svc 0
	ret

	.section  __DATA,__data
hello:  .ascii "Hej\n"
_mem: .word 0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
	`
