package main
import(
	"fmt"
	"time"
) 

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
	go say("hello")
	say("world")
	s := []int{1,2,3,4,5}
	c := make(chan int)
	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:],c)
	x,y := <-c,<-c
	fmt.Println(len(s),len(s)/2,x,y,x+y)
	test_chan := make(chan int ,2)
	test_chan <- 1
	test_chan <- 2
	fmt.Println(<-test_chan)
	fmt.Println(<-test_chan)
}
