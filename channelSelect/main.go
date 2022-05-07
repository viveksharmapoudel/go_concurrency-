package main

import (
	"fmt"
	"time"
)




func main(){

	ch1:= make(chan int )
	ch2 :=make(chan string )
	go passValue(ch1, ch2)

	for {
		select {
		case p:= <-ch1:
				fmt.Println(p)
				
		case q := <-ch2:
			fmt.Println(q)
		}
	}

}


func passValue(ch1 chan int , ch2 chan string, ){
	for i :=0; ; i++{
		ch1 <- i
		time.Sleep(time.Second)
		ch2 <- fmt.Sprintf("value is %d", i)
		time.Sleep(time.Second)
	}
}