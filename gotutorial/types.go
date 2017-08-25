package main
import (
	"fmt"
	"time"
	"math/rand"
	"math/cmplx"
	"math"
)
var c, python, java bool = true,false,true
var (
	love = true
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5+12i)
)


const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func add(x int,y int) int{
	return x+y
}

func swap(x, y string) (string,string){
	return y,x
} 

func show2(sum int)(x,y int) {
	x = sum % 2
	y = sum / 2 
	return 
}

func main(){
	fmt.Println("this is a test")
	// time
	fmt.Println("now the time is",time.Now(),)
	//random nubmer
	fmt.Println("there is random number",rand.Intn(30))
	// my favourite nubmer
	fmt.Println("my favourite nubmer is",math.Pi)
	// square root
	fmt.Printf("the square root of 13 is %g \n",math.Sqrt(13),)
	// test function
	fmt.Println("3 + 5 = ",add(3,5))
	// test mutiple results
	fmt.Println(swap("miss","you"))
	// test named return values
	fmt.Println(show2(13))
	// test varibale and type
	var i int = 49
	k :=88
	fmt.Println(i,k,c,java,python)
	fmt.Printf("Type: %T Value: %v\n", love, love)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	// test type conversion
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	// test type inference
	m :=  3
	fmt.Printf("Type: %T Value: %v\n", m, m)
	n := "test" 
	fmt.Printf("Type: %T Value: %v\n", n, n)
	// constant 
	const e = 2.718
	fmt.Printf("Type: %T Value: %v\n", e, e)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}













