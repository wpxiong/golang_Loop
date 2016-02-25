package main

import (
	"fmt"
)


type  LoopCompute struct {
   Result interface{}
   Init_value interface{}
} 


func (co *LoopCompute) compute(step interface{}){
   co.Result = co.Result.(int)  + step.(int)
}

func (co *LoopCompute) Init(){
   co.Init_value = interface{}(1)
   co.Result = interface{}(1)
}


func sendChan(ch chan int, cout int){
   if cout == 1 {
      ch <- cout
      return 
   }else{
       ch <- cout
       sendChan(ch,cout-1)
   }
}

func receiveChan(inputchan chan int,meth *LoopCompute,finish chan int){
   rec := <- inputchan
   if rec == 1 {
      meth.compute(rec)
      finish <- 1
   }else {
     meth.compute(rec)
     receiveChan(inputchan,meth,finish)
   }
}

func loopFunc(count int, meth *LoopCompute ) {
   inputchan := make(chan int,count)
   finish := make(chan int)
   meth.Init()
   go sendChan(inputchan,count)
   go receiveChan(inputchan,meth,finish)
   end := <- finish
   if end == 1 {
     return 
   }else {
     fmt.Println("Error")
   }
}


func main() {
  fmt.Println("start")
  meth := &LoopCompute{}
  loopFunc(40,meth)
  fmt.Print("Result:")
  fmt.Println(meth.Result.(int))
  fmt.Println("end")
}


