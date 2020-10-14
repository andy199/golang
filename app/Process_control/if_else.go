package main
//12
//流程控制
import "fmt"

//流程控制有if和for，而switch和goto主要是为了简化代码、降低重复代码而生的结构，属于扩展类的流程控制。
/*
if else(分支结构)

if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}
 */
//1.最简单的if语句
func main()  {
	//1.最简单的if语句
	flag := true
	if flag {
		fmt.Println("flag=true")
	}
	
	age := 30
	if age > 20 {
		fmt.Println("成年人")
	}

	//2.if语句的另一种写法
	if age := 34; age > 20 {//局部变量
		fmt.Println("yes 成年人")
	}

	//3.探讨上面两种写法的区别

	age1 := 30//当前区域内的是全局变量
	if age > 20 {
		fmt.Println("成年人", age1)
	}
	fmt.Println(age1)

	//4.输入成绩，如果成绩大于等于90输出A，如果小于90大于75输出B,否则输出C
    //全局作用域
	var score = 90
	if score >= 90 {
		fmt.Println("A")
	}else if score > 75{
		fmt.Println("B")
	}else {
		fmt.Println("C")
	}

	//局部作用域
	if score1 := 80; score1 >= 90{
		fmt.Println("a")
	}else if score1 > 75{
		fmt.Println("b")
	}else{
		fmt.Println("c")
	}
	
	//5.求两个数的最大值
	var a = 34
	var b = 89
	var max int
	if a > b {
		max = a
	}else {
		max = b
	}
	fmt.Println("a和b的最大值是", max)

}
