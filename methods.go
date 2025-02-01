package main

import "fmt"

type rect struct {
	width, height int
}

// value receiver gives a *rect the function area (r.area())
func (r *rect) area() int {
	return r.width * r.height
}

// you can also do it for rect

// You may want to use a pointer receiver type to avoid copying on
// method calls or to allow the method to mutate the receiving struct
func (r rect) perim() int {
	return 2 * r.width + 2 * r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area()) // acts over r's pointer
	fmt.Println("perim:", r.perim()) // acts over r


}
