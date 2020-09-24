//2
//变量
package main

import (
	"fmt"
	"math"
)

/*func main() {
	var age int 
	fmt.Println("My age is", age)
}*/
//go语言中变量声明了必须要使用，相同变量名不能重复使用
func main()  {
	var age int ///声明变量类型
	fmt.Println("my age is", age)
	age = 18
	fmt.Println("my age is", age)
	age = 80
	fmt.Println("my age is", age)
	///初始值变量
	var i int = 30
	fmt.Println("my age is", i)
	///变量类型推断
	var f  = 90
	fmt.Println("my age is", f)
	//多变量声明
	//var width, height = 100, 50
	//fmt.Println("width is", width, "height is", height)
	var width, height int 
	fmt.Println("width is", width, "height is", height)
	width = 100
	height = 50
	fmt.Println("width is", width, "height is", height)
	//不同类型的变量声明
	var (
		name = "andy"
		ages = 28
		heights int
	) 
	fmt.Println("my name is", name, ", age is", ages, "and height is", heights)
	//简洁变量声明，适用于局部变量声明，不能用于全局变量
	count := 10
	fmt.Println("Count =", count)
	//简洁方式声明多个变量
	names, agea := "tom", 30
	fmt.Println("my name is", names, "ags is", agea)


	a, b := 20, 30
	fmt.Println("a is", a, "b, is", b)
	b, c := 40, 50
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90
	fmt.Println("changed b is", b, "c is", c)

	d, e := 145.8, 543.8
	dd := math.Min(d, e)
	fmt.Println("Mini mun value is", dd)
}