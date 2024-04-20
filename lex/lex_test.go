package lex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLex(t *testing.T) {
	prog, err := Lex([]rune(">\n"))
	require.NoError(t, err)
	require.Len(t, prog, 1)
}
