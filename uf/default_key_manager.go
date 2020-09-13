package uf

// DefaultKeyManager manages key in the form of number
type DefaultKeyManager struct {
	defaultKey int
	lastKey    int
}

// GetUniqueKey generates a unique key in the form of number and updates the last key
func (km *DefaultKeyManager) GetUniqueKey() int {
	km.lastKey++
	return km.lastKey
}

// GetLastKey returns the last key returned to the caller
func (km *DefaultKeyManager) GetLastKey() int {
	return km.lastKey
}

// GetDefaultKey returns the default key that's never changed
func (km *DefaultKeyManager) GetDefaultKey() int {
	return km.defaultKey
}
