package main

import (
	"fmt"
	"time"
)


func main(){
    go process("order")
     process("refund")
}


func process(name string){
    for i:= 0; i<=5;i++{
        time.Sleep(time.Second/2)
        fmt.Printf("process: %d : %s\n", i , name)
    }
}