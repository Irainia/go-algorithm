package uf

// UnionFind is a type that performs union-find algorithm
type UnionFind struct {
	pointStorage []int
	pointWeight  []int
}

// Union unions two number points to form un-directed connection
func (uf *UnionFind) Union(p, q int) {
	rootP := uf.findRoot(p)
	rootQ := uf.findRoot(q)
	if rootP != rootQ {
		if uf.pointWeight[rootP] >= uf.pointWeight[rootQ] {
			uf.pointStorage[q] = rootP
			uf.pointWeight[p] += uf.pointWeight[q]
		} else {
			uf.pointStorage[p] = rootQ
			uf.pointWeight[q] += uf.pointWeight[p]
		}
	}
}

// IsConnected checks whether two number points are connected somehow
func (uf *UnionFind) IsConnected(p, q int) bool {
	return uf.findRoot(p) == uf.findRoot(q)
}

func (uf *UnionFind) findRoot(i int) int {
	for uf.pointStorage[i] != i {
		uf.pointStorage[i] = uf.pointStorage[uf.pointStorage[i]]
		i = uf.pointStorage[i]
	}
	return i
}

// NewUnionFind initializes union-find
func NewUnionFind(maximumNumber int) *UnionFind {
	pointStorage := make([]int, maximumNumber)
	for i := 0; i < maximumNumber; i++ {
		pointStorage[i] = i
	}
	pointWeight := make([]int, maximumNumber)
	for i := 0; i < maximumNumber; i++ {
		pointWeight[i] = 1
	}
	return &UnionFind{
		pointStorage: pointStorage,
		pointWeight:  pointWeight,
	}
}
