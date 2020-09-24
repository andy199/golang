package main

import "fmt"

func main() {
	/*
	go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。
	注意：
		布尔类型变量默认值为false
		go语言中不允许将整型强制转换布尔型
		布尔型无法参与数值运算，也无法与其他类型进行转换
	*/

	var flag = true
	fmt.Printf("%v--%T\n", flag, flag)
	var b bool
	fmt.Printf("%v", b)//默认值
}
