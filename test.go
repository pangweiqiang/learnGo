package main

import "fmt"

func main() {
   var i int= 1
   b:=2
   fmt.Println("Hello, World!")
   fmt.Println(i+b)
   for k:=0 ;k<10;k++ {
      fmt.Println(k);
   }
   fmt.Println(test(3))
   dowhile()
}
func test(k int) string{
   if k > 0 {
      return ">0"
   }else{
      return "<0"
   }
}
/**
do while 循环
**/
func dowhile() {
   for i:=0;i<10;{
      i++
      fmt.Println("haha")
   }

   
}
func test1(){
   var t int = 1
   if t == 1 {

   }  else{
      
   } 
}