//4
//常量
package main

import "fmt"

const pi = 3.1415
const e = 2.7182

func main() {
	//声明常量
	const a = 50
	fmt.Println(a)
	//一组常量声明
	const (
		name = "andy"
		age = 50
		country = "Canada"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	//字符串常量，无类型常量
	//类型常量和非类型常量
	const m  = "Sam"
	var na = m
	fmt.Printf("Type %T value %v", na, na)

	//布尔常量
	//布尔常量与字符串常量没有区别。它们是两个无类型化的常量true和false。
	//const trueConst  = true
	//type myBool bool
	//var defaultBool  = trueConst //allowed
	//var customBool myBool  = trueConst //allowed
	//defaultBool = customBool //not allowed

	//数值常数
	//数字常数包括整数，浮点数和复数常数。数字常数有些细微之处。
	const aa  = 5
	var intVar int = aa
	var int32Var int32 = aa
	var float64Var float64 = aa
	var complex64Var complex64 = aa
	fmt.Println("\nintVar:", intVar, "int32Var:", int32Var, "float64Var:", float64Var, "complex64Var:", complex64Var)

	//数值表达式
	//数字常量可以在表达式中自由混合和匹配，并且仅当将它们分配给变量或在代码中需要类型的任何位置使用时，才需要类型。

	var y  = 5.9 / 8
	fmt.Printf("y`s type is %T and value is %v", y, y)
	fmt.Println("-------")

	// iota是golang语言的常量计算器，只能在常量的表达式中使用
	const (
		n1 = iota
		n2
		n3
		n4
	)
	fmt.Println(n1, n2, n3, n4)
}
