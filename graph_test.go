package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnionFind(t *testing.T) {
	tests := []struct {
		name           string
		graph          map[int][]int
		expectedUnions int
	}{
		{"simple graph", map[int][]int{0: {1}, 1: {0}}, 1},
		{"simple graph with two unions", map[int][]int{0: {1}, 1: {0}, 2: {2}}, 2},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				UF := NewUnionFind(len(test.graph))
				for source, targets := range test.graph {
					for _, target := range targets {
						UF.Union(source, target)
					}
				}
				assert.Equal(t, test.expectedUnions, UF.count, "Your skills at algo suck!")
			},
		)
	}
}
