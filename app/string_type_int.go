package main
//10
import (
	"fmt"
	"strconv"
)

//1.string类型转换成整型（int）

func main()  {
	str := "123456"
	fmt.Printf("值：%v 类型：%T\n", str, str)//string
	
	/*
		ParseInt
		参数1：string数据
		参数2：进制
		参数3：位数 64 32 16
	*/
	//parseInt,有两个值
	//err 接收
	//_ 匿名忽略
	num, err := strconv.ParseInt(str, 10, 64)
	fmt.Printf("值：%v 类型：%T\n", num, num, err)

	num1, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("\n值：%v 类型：%T\n", num1, num1)

	/*
		ParseFloat
		参数1：string数据
		参数2：位数 64 32
	*/
	//parseFloat,有两个值
	//err 接收
	//_ 匿名忽略
	str1 := "123.56"
	num2, err := strconv.ParseFloat(str1, 64)
	fmt.Printf("值：%v 类型：%T\n", num2, num2, err)
	
	num3, _ := strconv.ParseFloat(str1, 64)
	fmt.Printf("\n值：%v 类型：%T\n", num3, num3)

	//5.数值类型没法和bool类型进行转换
}
