package main

import (
	"fmt"
	"sync"
)


func printNumber(num int64, wg *sync.WaitGroup){
	fmt.Println("Number: ", num)
	wg.Done()
}

func main(){

	var wg sync.WaitGroup

	wg.Add(2)
	go printNumber(3, &wg)
	go printNumber(10, &wg)
	wg.Wait()
}