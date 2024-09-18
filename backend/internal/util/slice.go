package util

func Remove[T any](items []T, filter func(item T) bool) []T {
	for i, item := range items {
		if filter(item) {
			return append(items[:i], items[i+1:]...)
		}
	}
	return items
}
