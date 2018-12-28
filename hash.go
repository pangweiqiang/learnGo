package main
import(
	"fmt"
	"unsafe"
)
func main()  {
	//hash1()
	//demo()
	//getKeyValArr()
	arrDemo()
}
func hash1(){
	var h map[int]string = map[int]string{
		1:"abc",
		2:"bc",
	}
	h[4] = "bac"
	delete(h,4)
	fmt.Println(h,len(h))
	var val,ok = h[4]
	fmt.Println(val,ok)
	if(ok){
		fmt.Println(val)
	}else{
		fmt.Println("not exist")
	}
}
//map 循环赋值  循环遍历
func demo(){
	var h = make(map[int]int,10);
	fmt.Println(len(h))
	for i:=0;i<5;i++{
		h[i] = i+1
	}
	fmt.Println(len(h))
	for j :=range h{
		fmt.Println(j)
	}
}
//获取map的 key val 数组
func getKeyValArr(){
	var h = make(map[string]int)
	h["a"] = 1
	h["b"] = 2
	h["c"] = 3
	fmt.Println(h,len(h))
	var kArr = make([]string,0,len(h))
	var vArr = make([]int,0,len(h))
	for k,v:=range h{
		kArr = append(kArr,k)
		vArr = append(vArr,v)
	}
	fmt.Println(kArr,vArr)
}
func arrDemo(){
	var arr = make([]int,10)
	var h = make(map[string]int)
	h["a"] = 1
	fmt.Println(arr,len(arr),cap(arr))
	fmt.Println(unsafe.Sizeof(arr))
}


