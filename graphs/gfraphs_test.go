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

func TestFindConnected(t *testing.T) {
	is := require.New(t)
	input := map[string][]string{
		"3": {},
		"4": {"6"},
		"6": {"4", "5", "7", "8"},
		"8": {"6"},
		"7": {"6"},
		"5": {"6"},
		"1": {"2"},
		"2": {"1"},
	}

	is.Equal(uint(3), get_component_count(input))
	is.Equal(uint(5), get_component_max_count(input))
}

func TestGrid(t *testing.T) {
	is := require.New(t)
	grid := [][]string{
		{"W", "L", "W", "W", "W"},
		{"W", "L", "W", "L", "W"},
		{"W", "W", "W", "L", "W"},
		{"L", "W", "L", "L", "W"},
		{"L", "L", "W", "W", "W"},
	}
	is.Equal(3, count_island(grid))
	is.Equal(2, minimum_island(grid))

}
