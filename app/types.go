//4
package main

import "fmt"

/*
func main() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	//a and b =false
	c := a && b
	fmt.Println("c:", c)
	//a or b =true
	d := a || b
	fmt.Println("d:", d)
}
*/

//有符号整数
/*
int8：表示8位带符号整数   大小： 8位    范围： -128至127
int16：表示16位有符号整数 大小： 16位 范围： -32768至32767
int32：表示32位有符号整数 大小： 32位 范围： -2147483648至2147483647
int64：表示64位有符号整数 大小： 64位 范围： -9223372036854775808至9223372036854775807
int：表示32或64位整数，具体取决于基础平台。除非需要使用特定大小的整数，否则通常应该使用int表示整数。
大小：在32位系统中为32位，在64位系统中为64位。
范围：在32位系统中为-2147483648至2147483647，在64位系统中为-9223372036854775808至9223372036854775807
*/
/*
func main()  {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))//64位系统输出8，32位输出4
	fmt.Printf("\ntype of b is %T, size od b is %d", b, unsafe.Sizeof(b))//64位系统输出8，32位输出4
}
*/
//无符号整数
/*
uint8：表示8位无符号整数	大小： 8位	范围： 0至255
uint16：表示16位无符号整数	大小： 16位	范围： 0到65535
uint32：表示32位无符号整数	大小： 32位	范围： 0至4294967295
uint64：表示64位无符号整数	大小： 64位	范围： 0至18446744073709551615
uint：表示32或64位无符号整数，具体取决于基础平台。
大小： 32位系统中为32位，64位系统中为64位。
范围：在32位系统中为0至4294967295，在64位系统中为0至18446744073709551615
*/

//浮点类型
//float32： 32位浮点数
//float64： 64位浮点数
/*
func main()  {
	a, b := 5.67, 8.98
	fmt.Printf("type of a %T b %T\n", a, b)//float64
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1 + no2, "diff", no1 - no2)
}
*/

//复数类型
//complex64：具有float32实部和虚部的
//复数complex128：具有float64实部和虚部的复数
/*
func main()  {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)

	cmul := c1 * c2
	fmt.Println("product:", cmul)
}
*/

//其他数值类型
//字节是uint8 的别名
//符文是int32的别名




//类型转换
func main()  {
	i := 55 //int
	j := 67.8 //float64
	//sum := i + j //int+float64 not allowed
	//类型不一 err
	//fmt.Println(sum)
	
	sum := i + int(j)//float64 转换 int
	fmt.Println(sum)

	var c float64 = float64(i)
	fmt.Println("c", c)
	
	sumf := float64(i) + j //int 转float64
	fmt.Println(sumf)
}

