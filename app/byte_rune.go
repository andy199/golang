package main
//7
//字符
//go语言的字符有两种
	//uint8类型 或者叫byte型，代表ASCLL码的一个字符，rune类型，代表一个UTF-8字符
import "fmt"

func main() {
	//go 中定义字符 字符属于int类型
	var a  = 'a'
	fmt.Printf("值：%v 类型:%T\n", a, a)//值：97 类型:int32
	//当我们直接输出byte（字符）的时候输出的是这个字符对应的码值 （ascll码）
	
	//原样输出字符
	var b = 'b'
	fmt.Printf("值：%c 类型:%T\n", b, b)//值：b 类型:int32
	
	//定义一个字符串输出字符串里面的字符
	var str = "this"
	fmt.Printf("值：%v 原样输出：%c 类型：%T\n", str[2], str[2], str[2])//值：105 原样输出：i 类型：uint8

	//一个汉字占用3个字节（utf-8），一个字母占用一个字节
	var str1  = "你好go"
	fmt.Println(len(str1))//8字节


	//定义一个字符 字符的值是汉字
	//golang 中汉字使用的是utf-8编码 编码后的值就是int类型
	var c  = '国'
	fmt.Printf("值：%v 类型：%T\n", c, c)//Unicode编码10进制 值：22269 类型：int32
	fmt.Printf("值：%c 类型：%T\n", c, c)//原样输出 值：国 类型：int32


	//通过循环输出字符串里面的字符
	s := "你好 golang"
	for _, v := range s {//rune
		fmt.Printf("%v(%c)", v, v)	//20320(你)22909(好)32( )103(g)111(o)108(l)97(a)110(n)103(g)
	}
	fmt.Println("\n=======")

	//修改字符串
	s1 := "big"
	byteStr := []byte(s1)
	byteStr[0] = 'p'
	fmt.Println(string(byteStr))

	s2 := "你好golang"
	runeStr := []rune(s2)
	runeStr[0] = '大'
	fmt.Println(string(runeStr))
}