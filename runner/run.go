package runner

import (
	"fmt"

	"bfc/lex"
	"bfc/parse"
)

type state struct {
	ip  int
	sp  byte
	mem [256]byte
}

func Run(ins []parse.Instruction) {
	s := state{}
	for s.ip < len(ins) {
		op := ins[s.ip]
		switch op.Op {
		case lex.Add:
			s.mem[s.sp]++
		case lex.Sub:
			s.mem[s.sp]--
		case lex.Left:
			s.sp--
		case lex.Righ:
			s.sp++
		case lex.Print:
			c := s.mem[s.sp]
			if c < 32 || c > 126 {
				fmt.Printf("(%d)", c)
			} else {
				fmt.Printf("%c", c)
			}
		case lex.JumpTo:
		case lex.JumpBack:
			if s.mem[s.sp] != 0 {
				s.ip = op.V
			}
		}

		s.ip++
	}

	fmt.Printf("\nState:\n%s", s)
}

func (s state) String() string {
	str := fmt.Sprintf("{\n\tip: %d\n\tsp: %d\n}", s.ip, s.sp)
	for i, b := range s.mem[:8] {
		if i%8 == 0 {
			str += fmt.Sprintf("\n")
		}
		if byte(i) == s.sp {
			str += fmt.Sprintf("->%d<-, ", b)
		} else {
			str += fmt.Sprintf("%d, ", b)
		}
	}

	return str
}
