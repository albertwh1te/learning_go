package main

import (
	"fmt"
	"./tree"
)
func Walk(t *tree.Tree, ch chan int){
	if t == nil{
		return
	}else {
		ch <- t.Value
		Walk(t.Left,ch)
		Walk(t.Right,ch)
	}
}

func Walker(t *tree.Tree) chan int{
	var ch chan int
	go func(){
		Walk(t,ch)
		// close(ch)
	}()
	return ch
}

func Same(t1, t2 *tree.Tree) bool{
	ch1,ch2 := Walker(t1),Walker(t2)
	for {
		v1,ok1 := <- ch1
		v2,ok2 := <- ch2
		if !ok1 || !ok2{
			return ok1 == ok2 
	}
		if v1 != v2{
			break
		}
}
	return false
}

func main(){
	// ch := make(chan int,10)
	// Walk(tree.New(1), ch)
	// close(ch)
	// for t := range ch{
	// fmt.Println(t)
	// }
	flag := Same(tree.New(1), tree.New(2))
	fmt.Println(flag)
}

