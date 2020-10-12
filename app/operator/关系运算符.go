package main

import "fmt"

//2
//关系运算符

/*
== 检查两个值是否相等，如果相等返回True 否则返回False。
!= 检查两个值是否不相等，如果不相等返回True 否则返回False。
>  检查左边值是否大于右边值，如果是返回True 否则返回False。
>= 检查左边值是否大于等于右边值，如果是返回True 否则返回False。
<  检查左边值是否小于右边值，如果是返回True 否则返回False。
<= 检查左边值是否小于等于右边值，如果是返回True 否则返回False。
*/

func main() {
//示例关系运算符的使用

var a1 = 9
var a2 = 8
fmt.Println(a1 == a2)//false
fmt.Println(a1 != a2)//true
fmt.Println(a1 > a2)//true
fmt.Println(a1 >= a2)//true
fmt.Println(a1 < a2)//false
fmt.Println(a1 <= a2)//false

flag := a1 > a2
if flag{
	fmt.Println("a1>a2")
}

if a1 == a2{
	fmt.Println("a1==a2")
}else {
	fmt.Println("a1!=a2")
}

}