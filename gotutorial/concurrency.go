package main
import(
	"fmt"
	"time"
) 

// func fibonacci_select(c,quit chan int){
// 	x,y := 0,1
// 	for {
// 		select {
// 			case c<-x:
// 			x,y = y,x+y
// 			fmt.Println(x,y)
// 			case <-quit:
// 			fmt.Println("quit")
// 			return 
// 		}

// 	}
// }

func fibonacci_select(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fibonacci(n int,c chan int){
	x,y :=  0,1
	for i := 0;i<n;i++{
		c <- x
		x,y = y,x+y
	}
	close(c)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func say(s string){
	for i:= 0;i<5;i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main(){
	// go say("hello")
	// say("world")
	// s := []int{1,2,3,4,5}
	// c := make(chan int)
	// go sum(s[:len(s)/2],c)
	// go sum(s[len(s)/2:],c)
	// x,y := <-c,<-c
	// fmt.Println(len(s),len(s)/2,x,y,x+y)
	// test_chan := make(chan int ,2)
	// test_chan <- 1
	// test_chan <- 2
	// fmt.Println(<-test_chan)
	// fmt.Println(<-test_chan)
	// t := make(chan int,10)
	// go fibonacci(cap(t),t)
	// for i := range t{
	// 	fmt.Println(i)
	// }

	// select with channel
	c := make(chan int)
	quit := make(chan int)
	go func(){
		for i := 0;i < 10;i++{
			fmt.Println(<-c)
		}
		quit <-0
	}()
	fibonacci_select(c,quit)
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(25 * time.Millisecond)
		}
	}
}


