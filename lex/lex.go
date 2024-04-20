package lex

type Token struct {
	Row, Col int
	Op       Op
}

type Op = byte

const (
	Add Op = iota
	Sub
	Righ
	Left
	JumpBack
	JumpTo
	Print
	Read
)

func Lex(bs []rune) ([]Token, error) {
	ts := []Token{}
	col := 0
	row := 1
	for _, r := range bs {
		col++
		t := Token{Col: col, Row: row}

		switch r {
		case '+':
			t.Op = Add
		case '-':
			t.Op = Sub
		case '>':
			t.Op = Righ
		case '<':
			t.Op = Left
		case '[':
			t.Op = JumpTo
		case ']':
			t.Op = JumpBack
		case '.':
			t.Op = Print
		case ',':
			t.Op = Read
		case '\n':
			row++
			col = 0
			fallthrough
		default:
			continue
		}

		ts = append(ts, t)
	}

	return ts, nil
}
