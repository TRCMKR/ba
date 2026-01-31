package utils

func SliceToSlice[A any, B any](slice []A, conv func(A) B) []B {
	newSlice := make([]B, 0, len(slice))

	for _, v := range slice {
		newSlice = append(newSlice, conv(v))
	}

	return newSlice
}

func SliceToMap[A any, K comparable, V any](slice []A, conv func(A) (K, V)) map[K]V {
	newMap := make(map[K]V, len(slice))

	for _, elem := range slice {
		k, v := conv(elem)
		newMap[k] = v
	}

	return newMap
}

func MapToSlice[K comparable, V any, A any](m map[K]V, conv func(k K, v V) A) []A {
	newA := make([]A, 0, len(m))

	for k, v := range m {
		newA = append(newA, conv(k, v))
	}

	return newA
}

// SliceDiff returns slice of elements from targetSlice which are not present in checkSlice
func SliceDiff[A comparable](targetSlice, checkSlice []A) []A {
	checkMap := make(map[A]struct{}, len(checkSlice))
	for _, v := range checkSlice {
		checkMap[v] = struct{}{}
	}

	diffSlice := make([]A, 0)
	for _, v := range targetSlice {
		if _, found := checkMap[v]; !found {
			diffSlice = append(diffSlice, v)
		}
	}

	return diffSlice
}

// SliceFullDiff returns slice of different elements from sliceA which are not in slice B
// and vice versa
func SliceFullDiff[A comparable](sliceA, sliceB []A) []A {
	diffSlice := make([]A, 0)

	diffSlice = append(diffSlice, SliceDiff(sliceA, sliceB)...)
	diffSlice = append(diffSlice, SliceDiff(sliceB, sliceA)...)

	return diffSlice
}
