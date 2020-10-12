package main

import "fmt"

//1
/*
运算符
	1.算术运算符
	+ 相加
	- 相减
	* 相乘
	/ 相除
	% 求余=被除数 -（被除数/除数）* 除数
	1.除法注意：如果运算符的数都是整数，那么除后，去掉小数部分，保留整数部分
	2.取余注意：余数=被除数 -（被除数/除数）* 除数
	3.注意：++（自增），和，--（自减）在go语言中是单独的语句，并不是运算符。
*/

func main()  {
	var a = 6
	var b = 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	
	var c = a * b
	fmt.Println(c)

    //1.除法注意：如果运算符的数都是整数，那么除后，去掉小数部分，保留整数部分
    
    var e = 10
    var f = 3
    fmt.Println(e / f)//3
    
    var h = 10.0
    var i = 3.0
    fmt.Println(h / i)//3.3333333333333335

    //2.取余注意：余数=被除数 -（被除数/除数）* 除数
    var z = 10
    var y = 3
    fmt.Println(z % y)//1

    var v = -10
    var n = 3
    fmt.Println(v % n)// -1 |-10-(-10/3)*3=-1
    fmt.Println(10 % -3)//1 |10-(10/-3)*-3=1

    //3.注意：++（自增），和，--（自减）在go语言中是单独的语句，并不是运算符。
    //1.注意：在golang中，++和-- 只能独立使用，错误写法如下：
   	/*
   	var j int = 8
    var k int
    k = j++ //错误，j++只能独立使用
    k = j-- //错误，j--只能独立使用
    fmt.Println(k)
   	*/

   	//2.注意：在golang中没有前++ --错误写法如下：
   	/*
   	var m = 12
   	++m //错误写法
   	fmt.Println(m)
   	*/

   	//3.正确写法，单独使用
	var s = 12
	s++
	fmt.Println(s)
	s--
	fmt.Println(s)
}