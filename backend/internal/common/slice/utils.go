package slice

func ContainsFiltered[T any](items []T, compare func(T) bool) bool {
	for _, v := range items {
		if compare(v) {
			return true
		}
	}
	return false
}

func ConvertToPointerSlice[T any](items []T) []*T {
	var pointer []*T
	for i := range items {
		pointer = append(pointer, &items[i])
	}
	return pointer
}

func MapSlice[T any, U any](input []T, mapper func(item T) U) []U {
	output := make([]U, len(input))
	for i, v := range input {
		output[i] = mapper(v)
	}
	return output
}
