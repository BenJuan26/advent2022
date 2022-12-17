package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReduceAll(t *testing.T) {
	ranges := Ranges([]Range{{0, 5}, {7, 10}, {12, 19}})
	ranges = ranges.ReduceAll()
	require.Equal(t, Ranges([]Range{{0, 5}, {7, 10}, {12, 19}}), ranges)

	ranges = ranges.Push(Range{4, 7})
	ranges = ranges.ReduceAll()
	require.Equal(t, Ranges([]Range{{0, 10}, {12, 19}}), ranges)
}
