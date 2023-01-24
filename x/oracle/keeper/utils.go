package keeper

// contains returns true if x is in ls
func contains[T comparable](x T, ls []T) bool {
	for i := range ls {
		if ls[i] == x {
			return true
		}
	}

	return false
}
