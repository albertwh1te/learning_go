package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x,n,lim float64) float64 {
	if v := math.Pow(x,n);v<lim{
		return v
	} else {
		fmt.Printf("%g => %g",v,lim)
	}
	return lim
}

func newton_method(x float64,n int) float64{
	z := 1.0
	for i :=0 ;i <n;i++ {
		z = z - (z*z - x)/(2*z)
	}
	return z
}


func showsystem(){
	fmt.Println("GO runs on")
	switch os:= runtime.GOOS;os{
		case "darwin":
		fmt.Println("OS X")
		case "linux":
		fmt.Println("Linux")
		default:
		fmt.Printf("%s",os)
	}
}


func showtime(){
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func defercount(){
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}


func main() {
	// basic for loop
	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}
	fmt.Println(sum1)
	sum2 := 1
	// for is go'while
	for sum2 < 1000 {
		fmt.Println(sum2)
		sum2 += sum2
	}
	fmt.Println(sum2)
	// basic if 
	fmt.Println(sqrt(2),sqrt(-9))
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
	fmt.Println(newton_method(3,100),sqrt(3))
	// switch syntax
	showsystem()
	// switch without no condition
	showtime()	
	// defer test
	defer fmt.Println("world")
	fmt.Println("hello")
	defercount()
}
