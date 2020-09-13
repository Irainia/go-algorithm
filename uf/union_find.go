package uf

// KeyManager manages key
type KeyManager interface {
	GetUniqueKey() int
	GetLastKey() int
	GetDefaultKey() int
}

// UnionFind is a type that performs union-find algorithm
type UnionFind struct {
	pointStorage []int
	keyManager   KeyManager
}

// Union unions two number points to form un-directed connection
func (uf *UnionFind) Union(p, q int) {
	newKey := uf.keyManager.GetUniqueKey()
	if uf.pointStorage[p] == uf.pointStorage[q] {
		if uf.pointStorage[p] == uf.keyManager.GetDefaultKey() {
			uf.pointStorage[p] = newKey
			uf.pointStorage[q] = newKey
		}
	} else {
		if uf.pointStorage[p] == uf.keyManager.GetDefaultKey() {
			uf.pointStorage[p] = uf.pointStorage[q]
		} else if uf.pointStorage[q] == uf.keyManager.GetDefaultKey() {
			uf.pointStorage[q] = uf.pointStorage[p]
		} else {
			keyForP := uf.pointStorage[p]
			keyForQ := uf.pointStorage[q]
			for i := 0; i < len(uf.pointStorage); i++ {
				if uf.pointStorage[i] == keyForP ||
					uf.pointStorage[i] == keyForQ {
					uf.pointStorage[i] = newKey
				}
			}
		}
	}
}

// IsConnected checks whether two number points are connected somehow
func (uf *UnionFind) IsConnected(p, q int) bool {
	if uf.pointStorage[p] == uf.pointStorage[q] {
		return uf.pointStorage[p] != uf.keyManager.GetDefaultKey()
	}
	return false
}

// NewUnionFind initializes union-find
func NewUnionFind(maximumNumber int, keyManager KeyManager) *UnionFind {
	return &UnionFind{
		pointStorage: make([]int, maximumNumber),
		keyManager:   keyManager,
	}
}
