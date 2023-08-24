package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTraversals(t *testing.T) {
	is := require.New(t)
	input := map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	expectedOutputDFS := []string{
		"a", "c", "e", "b", "d", "f",
	}
	output := dfs_traversal(input, "a")
	for i, v := range expectedOutputDFS {
		if v != output[i] {
			t.Fatalf("Failure in DFS Expected output mismatch at %d expected %s got %s slice received %v", i, v, output[i], output)
		}
	}

	expectedOutputBFS := []string{
		"a", "b", "c", "d", "e", "f",
	}
	outputbfs := bfs_traversal(input, "a")
	for i, v := range expectedOutputBFS {
		if v != outputbfs[i] {
			t.Fatalf("Failure in DFS Expected output mismatch at %d expected %s got %s slice received %v", i, v, outputbfs[i], outputbfs)
		}
	}

	is.True(has_path(input, "a", "c"))
	is.False(has_path(input, "e", "f"))
	is.True(has_path(input, "a", "f"))
	is.False(has_path(input, "c", "f"))
	is.True(has_path(input, "a", "e"))

}
