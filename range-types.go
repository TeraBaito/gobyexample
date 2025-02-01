package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	// like enumerate (python)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	// iterate over key, value
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// iterate over key
	for k := range(kvs) {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
