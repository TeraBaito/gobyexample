package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// skip types until the end
func plusPlus(a, b, c int) int {
	return a + b + c
}

func test(a, b int, c, d string) {
	fmt.Println(a, b, c, d)
}

// multiple return values
func vals() (int, int) {
	return 3, 7
}

// variadic function
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)

	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)

	d, _ := vals()
	fmt.Println(d)

	test(1, 2, "h", "wow")

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
