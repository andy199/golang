package main

import "fmt"

//定义一个方法
func test()  bool{
	fmt.Println("test....")
	return true
}

//3
//逻辑运算符
func main()  {
	/*
	&& 逻辑 and 运算符，如果两边的操作数都是true，则为true，否则为false
	|| 逻辑 or 运算符， 如果两边的操作数有一个是true，则为true，否则为false
	！ 逻辑 not 运算符，如果条件为true，则为false，否则为true
	*/
	// && and
	var a = 23
	var b = 8
	fmt.Println(a > 10 && b < 10)//true
	fmt.Println(a > 24 && b < 10)//false
	fmt.Println(a > 5 && b < 6) //false

	// || or
	var c = 23
	var d = 8
	fmt.Println(c > 10 || d < 10)//true
	fmt.Println(c > 24 || d < 10)//true
	fmt.Println(c > 5 || d < 6)//true
	fmt.Println(c == 5 || d < 6)//false
	
	// not 取反
	flag := true
	fmt.Println(!flag)//false
	flag1 := false
	fmt.Println(!flag1)//true

	//逻辑与和逻辑或短路
	/*
	输出
	test....
	执行
	*/
	var x = 10
	if x > 9 && test() {
		fmt.Println("执行")
	}
	/*
	输出
	空，短路
	*/
	var v = 10
	if v > 11 && test() {
		fmt.Println("执行")
	}

	/*
	输出
	test....
	执行
	*/
	var y = 10
	if y > 11 || test() {
		fmt.Println("执行")
	}

	/*
		输出
		空 ，短路
		执行
	*/
	var k = 10
	if k < 11 || test() {
		fmt.Println("执行")
	}
}