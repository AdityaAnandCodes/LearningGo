package main

import (
	"fmt"
)


func getStringLength(s string, ch chan int){
	ch <- len(s)
}

func main(){
	ch := make(chan int)
	go getStringLength("Ramu Lal" , ch)
	
	result := <-ch 
	fmt.Println(result)
}