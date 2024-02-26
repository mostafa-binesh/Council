package utils

// difference returns a new slice containing elements that are in sliceA but not in sliceB.
// It works with slices of any comparable type.
func Difference[T comparable](sliceA, sliceB []T) []T {
	m := make(map[T]bool)
	var diff []T

	// Mark all elements of sliceB in the map.
	for _, item := range sliceB {
		m[item] = true
	}

	// Add elements to diff slice if they're not marked in the map.
	for _, item := range sliceA {
		if _, found := m[item]; !found {
			diff = append(diff, item)
		}
	}

	return diff
}
