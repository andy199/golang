package main
//9
import (
	"fmt"
	"strconv"
)

//其他类型转换成string类型
//注意：sprintf使用中需要注意转换的格式 int 为%d, float 为%f, bool 为%t, byte 为%c


func main() {

	var a int = 20
	var b float64 = 12.456
	var c bool = true
	var d byte = 'a'

	str1 := fmt.Sprintf("%d", a)
	fmt.Printf("值：%v 类型：%T\n", str1, str1)

	str2 := fmt.Sprintf("%.2f", b)
	fmt.Printf("值：%v 类型：%T\n", str2, str2)

	str3 := fmt.Sprintf("%t", c)
	fmt.Printf("值：%v 类型：%T\n", str3, str3)
	
	str4 := fmt.Sprintf("%c", d)
	fmt.Printf("值：%v 类型：%T\n", str4, str4)


	//2.通过strconv 把其他类型转换成string类型
	/*
		FormatInt
		参数1: int64的数值
	 	参数2：传值int类型的进制 （数字）
	*/
	var e int = 20
	str5 := strconv.FormatInt(int64(e), 10)
	fmt.Printf("值：%v 类型：%T\n", str5, str5)

	/*
		参数 1：要转换的值
		参数 2：格式化类型 ‘f’ (-ddd.dddd)，
				'b'(-ddddp±ddd, 指数为二进制)，
				'e'（-d.dddde±dd, 十进制指数），
				'E'（-d.ddddE±dd, 十进制指数），
				'g'（指数很大时用'e'格式，否则'f'格式），
				'G' （指数很大时用'E'格式，否则'f'格式）
		参数 3：保留的小数点 -1 （不对小数点格式化）
		参数 4：格式化的类型 传入 64 32
	*/
	var f float32 = 20.23334
	str6 := strconv.FormatFloat(float64(f), 'f', 3, 64)
	fmt.Printf("值：%v 类型：%T\n", str6, str6)
	
	//3.布尔转换string //了解，没有任何意义
	str7 := strconv.FormatBool(true)
	fmt.Printf("值: %v 类型：%T\n", str7, str7)

	//4.字符转换string //了解，意义不大
	g := 'b'
	str8 := strconv.FormatUint(uint64(g), 10)
	fmt.Printf("值：%v 类型：%T\n", str8, str8)
}