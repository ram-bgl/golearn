package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	modifyValue(r)

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	modifyRef(rp)
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	fmt.Printf("Address of r: %p\n", rp)
}

func modifyRef(rp *rect) {
	rp.width = 20
	rp.height = 10
	fmt.Printf("Address of modifyRef var: %p\n", rp)
}

func modifyValue(r rect) {
	r.width = 20
	r.height = 10
	fmt.Printf("Address of modifyValue var: %p\n", &r)
}
