package main

import (
	"fmt"
	"math"
)

func main() {
    fmt.Printf("max float32 = %+v\n", uint(math.MaxUint64))
	x := make([]int, 3)
	x[1]=1
	fmt.Println(x,"--", len(x),cap(x))
	x = append(x, 4)
	fmt.Println(x,"--", len(x),cap(x))
	// x[4]=4
	x = append(x, 4)
	fmt.Println(x,"--", len(x),cap(x))

	var a1 [10]int
	var a2 [] int
	var a3 [] int = make([]int, 0)
	var a4 [] int = make([]int, 2)
	fmt.Println(a1,a2, a3, a4)

	// a2[0]=5
	// a3[0]=5
	a4[0]=5

	a2 = append(a2, 1)
	a3 = append(a3, 2)
	a4 = append(a4, 4)
	fmt.Println(a1,a2, a3, a4)

	p1 := Person{Name: "Tiago", Age: 46, Info: Info{Address: "Rua1"}}
	// p1 := Person{Name: "Tiago", Age: 46, Data: Info{Address: "Rua1"}}
	fmt.Printf("%+v",p1)
	fmt.Printf("%+v",p1.Address)
	p1.Blocked=true

	var p *int=&a4[0]
	fmt.Println()
	fmt.Println(*p)

}

type Person struct{
	Name string
	Age int
	Info
	Blocked bool
}

type Info struct{
	Address string
}