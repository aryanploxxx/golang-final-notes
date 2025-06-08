package main

// #1
// func printSlice(items []int) {
// 	for _, item := range items {
// 		println(item)
// 	}
// }

// #2
// func printSliceString(items []string) {
// 	for _, item := range items {
// 		println(item)
// 	}
// }

/*
	type Num interface {
		int | float64 | float32 | int32 | int64 | uint32 | uint64 | uint | uint8 | uint16 | int8 | int16
	}
	func printSlice[T Num](items []T) {
*/

// T any can be replaced with T interface{} -> both do the same things
// If we want to scope it down to allow only 2 types, we can use T interface{int, string} or func printSlice[T int | string](items []T) {
// [T comparable] -> T can be any comparable type like int, string, bool, float32, float64, struct with comparable fields
// [T comparable, U string]
func printSlice[T any](items []T) {
	for _, item := range items {
		println(item)
	}
}

func main() {
	items := []int{1, 2, 3, 4, 5}
	// #1
	// printSlice(items)

	itemsString := []string{"a", "b", "c", "d", "e"}
	// #2
	// printSliceString(itemsString)

	printSlice(items)
	printSlice(itemsString)
}
