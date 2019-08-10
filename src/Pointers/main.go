package main

import (
	"fmt"
	"reflect"
)

func main() {
	per1 := person{age: 20}
	var per2 *person
	per2 = &per1
	fmt.Println("1. ****struct Initial value", per1.age)
	per2.age = 90
	fmt.Println("1. ****struct After changing value (without pointer)", per1.age)

	i := 0
	p1 := &i
	fmt.Println("1. ****Initial value", i)
	doSomethingWithoutPointer(i)
	*p1 = 90
	fmt.Println("1. ****After changing value (without pointer)", i)
	fmt.Println("&i", p1, reflect.TypeOf(p1), *p1)

	var p *int

	fmt.Println("\n2. ****Initial value", p)
	doSomethingWithPointer(p)
	fmt.Println("2. ****After changing value (with pointer)", p)

}

func doSomethingWithPointer(point *int) {
	*point = 10
}

func doSomethingWithoutPointer(point int) {
	point++
}

type person struct {
	age int
}
