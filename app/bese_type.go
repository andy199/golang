package main

import "fmt"
//8
//基本数据类型之间转换


func main() {
	//整型之间转换
	var a int8 = 20
	var b int16 = 40
	fmt.Println(int16(a) + b)

	//浮点型和整型之间转换
	var c float32 = 3.2
	var d int16 = 6
	var e = c + float32(d)
	fmt.Println(e)

	//浮点型之间转换
	var f float32 = 4.3
	var g float64 = 1.2
	var h = float64(f) + g
	fmt.Printf("%.2f\n", h)

	//整型和浮点型之间转换
	var i float32  = 20.23
	var j int = 40
	fmt.Println(i + float32(j))

	//注意：转换的时候建议从低位转换成高位，高位转换成低位的时候如果转换不成功就会溢出，和我们想的结果不一样。
	var k int8 = 20
	var l int16 = 140
	fmt.Println(int16(k) + l)

}
