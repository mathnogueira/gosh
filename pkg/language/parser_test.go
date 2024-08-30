package language_test

import (
	"context"
	"testing"

	"github.com/mathnogueira/gosh/pkg/language"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	parser := language.NewParser()
	x, err := parser.Parse(context.Background(), "../../examples/main/main.gosh")
	require.NoError(t, err)
	require.NotNil(t, x)
}
