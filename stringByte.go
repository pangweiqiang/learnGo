package main
import "fmt"
func main(){
	fmt.Println(len("庞"))
	test()
} 
func test(){
	var s = "中国china"
	for i:=0;i<len(s);i++{
		fmt.Printf("%x ",s[i]);
	}
}