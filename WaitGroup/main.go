package main

import (
	"fmt"
	"sync"
	"time"
)


func printNumber(num int64, wg *sync.WaitGroup){

	fmt.Println("Number: ", num ,"start")
	time.Sleep(time.Second* 8)
	fmt.Println("Number:", num, "finish")
	wg.Done()
}

func main(){

	var wg sync.WaitGroup

	wg.Add(2)
	go printNumber(3, &wg)
	go printNumber(10, &wg)
	wg.Wait()
}