package main

import (
	"fmt"
	"sync"
)

// CarState -> simple struct which denote state of the car
type CarState struct{
	moving bool 
	speed float64
	windowOpen bool 
	mutexState sync.RWMutex
}

func(c *CarState) setMoving(x bool){
	c.moving= x
}

func(c *CarState) setSpeed(spd float64){
	c.speed= spd
}

func(c *CarState) setWindowOpen(win bool){
	c.windowOpen=win
}

func(c *CarState) readCar(){
	c.mutexState.RLock()
	fmt.Print(c)
	c.mutexState.RUnlock()
}


func main(){

	car := CarState{
		moving: false,
		speed: 0,
		windowOpen: false,
	}
	car.readCar()

	go changeValue()

}

func changeValue(){


	for i:=0; ; i++{
		select{
		}
	}
}