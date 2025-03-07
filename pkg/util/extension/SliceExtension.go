package extension

func GetLastIndexFrom[T any](slice []T) int {
	if len(slice) == 0 { return 0 }
	return len(slice)-1
}