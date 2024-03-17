package main

import "fmt"

var maxSearch int
var subset []int

// searchSubsets creates subsequences from 1 to maxSearch, requires maxSearch and subset global vars.
// Start with call to 1.
// This function is destructive -- starts with empty slice and ends with it.
// Way to O(2ⁿ), do not use with maxSearch ≥ 20
func searchSubsets(k int) {
	if k == maxSearch+1 {
		fmt.Println(subset) //do something with subset
		return
	}

	subset = append(subset, k)
	searchSubsets(k + 1)
	subset = subset[:len(subset)-1]
	searchSubsets(k + 1)
}
