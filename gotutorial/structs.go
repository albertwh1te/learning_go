package main

import (
	"fmt"
)

func pointdemo(){
	i, j := 42, 2701
	
		p := &i         // point to i
		fmt.Println(*p) // read i through the pointer
		*p = 21         // set i through the pointer
		fmt.Println(i)  // see the new value of i
	
		p = &j         // point to j
		*p = *p / 37   // divide j through the pointer
		fmt.Println(j) // see the new value of j
}

type Vertex struct{
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main(){
	pointdemo()
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v)
	// test struct point
	p := &v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(v1, p, v2, v3)
}

