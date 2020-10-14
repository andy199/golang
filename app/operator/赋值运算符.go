package main

import "fmt"

//4
/*
=	简单的赋值运算符，将一个表达式的值赋给一个左值
+=	相加后再赋值
-=	相减后再赋值
*=	相乘后再赋值
/=	相除后再赋值
%=	求余后再赋值
<<=	左移后赋值
>>=	右移后赋值
&=	按位与后赋值
|=	按位或后赋值
^=	按位异或后赋值
*/
func main()  {
	var a = 20
	a = 21//赋值
	fmt.Println(a)
	
	var b = 21 + 3
	fmt.Println(b)
	
	var c = 11
	d := c
	fmt.Println(d)

	var e = 12
	f := e + 3
	fmt.Println(f)

	var g = 10
	g += 2
	fmt.Println(g)
	
	var l = 10
	l -= 1
	fmt.Println(l)
	
	var k = 10
	k *= 2
	fmt.Println(k)
	
	var v = 10
	v /= 3
	fmt.Println(v)//3

	var s float64 = 10
	s /= 3
	fmt.Println(s)//3.3333333333333335

	var n = 10
	n %= 3
	fmt.Println(n)//1
	
}