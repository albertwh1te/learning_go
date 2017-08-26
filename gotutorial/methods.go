package main 
import (
	"fmt"
	"math"
	)

type Person struct{
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte 

func (a IPAddr) String() string{
	return fmt.Sprintf("%v.%v.%v.%v",a[0],a[1],a[2],a[3])
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type T struct{
	S string 
}
type I interface{
	M()
}
func (t T) M(){
	fmt.Println(t.S)
}

type Myfloat float64
func (f Myfloat) fAbs() float64{
	if f < 0{
		return float64(-f)
	}
	return float64(f) 
}

type Vertex struct{
	X,Y float64
}

func(v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)  
}

// this is a methods it takes not pointer but also values
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// this is piont function it only takes pointer
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main(){
	fmt.Println("test")
	// v := Vertex{3, 4}
	// fmt.Println(v.Abs())
	// f := Myfloat(-math.Sqrt2)
	// fmt.Println(f.fAbs())
	// v.Scale(10)
	// fmt.Println(v.Abs())
	// ScaleFunc(&v,10)
	// fmt.Println(v.Abs())
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
	s := T{"impliect interface test"}
	s.M()

	// the empty interface
	// var i interface{}
	// describe(i)

	// i = 42
	// describe(i)

	// i = "hello"
	// describe(i)

	// Type assertions
	var i interface{} = "hello"
	m := i.(string)
	fmt.Println(m)
	m, ok := i.(string)
	fmt.Println(m, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)
	// f = i.(float64) // panic
	fmt.Println(f)
	do(42)
	do("hello")
	do(true)
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}