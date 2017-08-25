package main 
import (
	"fmt"
	"strings"
	"math"
)

func fibonacci() func() int {
	x,y  := 0,1
	return func() int{
		x,y = y,x+y
		return x
	}
}
func showfibonacci(){
	f := fibonacci()
	for i:=0;i<10;i++{
		fmt.Println(f())
	}
}

// functions are values too
func compute(fn func(float64,float64) float64,a float64,b float64) float64{
	return fn(a,b)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func showclosure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	s_list := strings.Fields(s)
	for _ ,i := range s_list{
		m[i] += 1
	}
	return m
}

func mapabout(){
	type Vertex struct{
		Lat,Long float64
	}
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
    m["google"] = Vertex{
		56.3, -14.7,
	}
	fmt.Println(m["Bell Labs"],m)
	mapcrud()
}

func mapcrud(){	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}


func showrange(pow []int){
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}	
}
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

func slicelenandcap(){
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func makeslice(){
	a := make([]int, 5)
	fmt.Println("a",a)
	b := make([]string,0,10)
	fmt.Println("b",b,len(b),cap(b))
}


func tictactoeboard(){
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}	
}

func sliceappend(){
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func main(){
	// pointer
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

	// arrays
	var a [2]string
	fmt.Println(a[0], a[1],a)
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	
	// slices
	// var s []int = primes[1:4]
	s := primes[1:4] //in function we can use :=  instead of var
	fmt.Println(s)
	s[2] = 4
	fmt.Println(s)
	fmt.Println(primes)
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	g := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}	
	fmt.Println(g)
	// slicelenandcap()
	makeslice()
	tictactoeboard()
	sliceappend()
	// range test
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	showrange(pow)
	row := make([]int, 10)
	for i := range row {
		row[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range row {
		fmt.Printf("%d\n", value)
	}
	// map 
	mapabout()
	// function as value
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	// fmt.Println(compute(math.Pow,3,4))
	fmt.Println(compute(hypot,3,4))
	// showclosure()
	showfibonacci()
}


