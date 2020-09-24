 //1
 //hi go
package main

import "fmt"


func main() {
	fmt.Print("hi", "A")
	fmt.Println( "hi", "B")
	fmt.Println( "hi golang")

	//print 和 println区别
	//一次输入多个值的时候println 中间有空格print没有，println会自动换行，print不会
	//printf 是格式化输出，在很多场景下比println更方便
	a := 10
	b := 20
	c := 30
	fmt.Println("a=", a, "b=", b, "c=", c) //a= 10,b= 20,c= 30
	fmt.Printf("a=%d, b=%d ,c=%d", a, b, c) //a=10,b=20,c=30
}
