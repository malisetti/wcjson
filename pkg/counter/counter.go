package counter

// KeyCount returns the total number of keys in all nested maps within v.
// Scalars and arrays contribute no keys; array elements are traversed recursively.
func KeyCount(v any) int {
	switch x := v.(type) {
	case map[string]any:
		n := len(x)
		for _, val := range x {
			n += KeyCount(val)
		}
		return n
	case []any:
		n := 0
		for _, val := range x {
			n += KeyCount(val)
		}
		return n
	default:
		return 0
	}
}

// LeafCount returns the number of non-container values within v.
// Maps and arrays are containers; scalars and null each count as one leaf.
func LeafCount(v any) int {
	switch x := v.(type) {
	case map[string]any:
		n := 0
		for _, val := range x {
			n += LeafCount(val)
		}
		return n
	case []any:
		n := 0
		for _, val := range x {
			n += LeafCount(val)
		}
		return n
	default:
		return 1
	}
}
