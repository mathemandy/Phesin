package main

func PrintSlice[T any](i []T) {
	for _, v := range i {
		print(v)
	}
}

func main() {
	PrintSlice([]int{1, 2, 3, 4})
}
