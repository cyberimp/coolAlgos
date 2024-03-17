package main

// UnionFind is structure for finding connectivity in graph, represented as whatever you like
type UnionFind struct {
	parent []int
	size   []int
	count  int
}

// NewUnionFind creates initialized UnionFind structure
func NewUnionFind(s int) UnionFind {
	parent := make([]int, s)
	size := make([]int, s)
	for i := 0; i < s; i++ {
		parent[i] = i
		size[i] = 1
	}
	return UnionFind{parent, size, s}
}

// Find the root of set
func (u *UnionFind) Find(q int) int {
	for q != u.parent[q] {
		q = u.parent[q]
	}
	return q
}

// Union two sets
func (u *UnionFind) Union(a, b int) {
	rootA := u.Find(a)
	rootB := u.Find(b)

	if rootA == rootB {
		return
	}

	if u.size[rootA] < u.size[rootB] {
		u.parent[rootA] = rootB
		u.size[rootB] += u.size[rootA]
	} else {
		u.parent[rootB] = rootA
		u.size[rootA] += u.size[rootB]
	}
	u.count--

	return
}
