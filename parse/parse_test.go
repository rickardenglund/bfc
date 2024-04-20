package parse

import (
	"testing"

	"github.com/stretchr/testify/require"

	"bfc/lex"
	"bfc/runner"
)

func TestParse(t *testing.T) {
	ts, err := lex.Lex([]rune(">"))
	require.NoError(t, err)

	prog, err := Parse(ts)
	require.NoError(t, err)
	require.Len(t, prog, 1)

	runner.Run(prog)

}
