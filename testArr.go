package main

import "fmt"
/**
func main() {
   var a[9] int
   var b = [5]int{1,2,3,4,5}
   fmt.Println(a)
   fmt.Println(b)
   //test()
   //arr_qiepian()
   //arr_qiege()
   //arr_kuorong()
   arr_copy()
}**/

func test(){
   var k [5] int = [5]int{1,2,3,4,5}
   //var p [5] int
   var p = k
   k[0] = 100
   fmt.Println(p,k)
   for i,v := range k{
      fmt.Println(i,v)
   }
   var n = len(k)
   for j:=0 ; j<n ;j++{
      fmt.Println(k[j])   
   }
   fmt.Println(p,k)
}
func arr_qiepian(){
   var arr []int = make([]int,5,10);//满容切片和普通数组区别 1 var arr []int = []int{1,23} 2 var arr [5] int 
   fmt.Println(cap(arr),len(arr))   
   var arr1 = arr
   arr1[5] = 100
   fmt.Println(arr,arr1)
}
//切片的切割
func arr_qiege(){
   var arr []int = []int{0,1,10,3,4,5}
   var a = arr[1:]
   fmt.Println(arr,len(arr),cap(arr),a,len(a),cap(a))
}
//切片的扩容
func arr_kuorong(){
   var arr []int = []int{1,2,3,4,5}
  var  arr1 []int = make([]int,1024)
  arr = append(arr,1)
  arr1 = append(arr1,1)
   fmt.Println(arr,len(arr),cap(arr),arr1,len(arr1),cap(arr1))
  
}
//copy的使用
func arr_copy(){
   var arr = make([]int ,5,10)
   l := len(arr)
   for i:=0;i<l;i++{
      arr[i] = i+1
   }
   var arr1 []int = make([]int,2,5)
   copy(arr1,arr)
   fmt.Println(arr1)
}