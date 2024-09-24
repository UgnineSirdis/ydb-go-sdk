package sugar

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStackRecord(t *testing.T) {
	require.Equal(t, `github.com/UgnineSirdis/ydb-go-sdk/v3/sugar.TestStackRecord(stack_test.go:10)`, StackRecord(0))
}
