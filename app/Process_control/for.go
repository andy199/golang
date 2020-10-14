package main

import "fmt"

//2
//for循环结构
//for格式
/*
for 初始语句;条件表达式;结束语句{
    循环体语句
}
 */

func main()  {
	//打印1-10数字
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	
	v := 0
	for ; v < 10; v++ {
		fmt.Println(v)
	}

	k := 0
	for k < 10 {
		fmt.Println(k)
		k++
	}

	//无限循环
	/*
	for {
		循环体语句
	}
	 */
	a := 1
	for {
		if a <= 10 {
			fmt.Println(a)
		}else {
			break//跳出循环
		}
		a++
	}

//练习
	//1.打印0-50所有的偶数
	for h := 0; h <= 50; h++ {//循环0-50
		if h % 2 == 0{//判断是不是被2整除
			fmt.Println(h)
		}
	}

	//2.求 1+2+3+4+...100的和
	/*
	循环开始
	初始 g := 1
	条件 g <= 100 小于100
	执行sum += g //sum +=sum+1
	计数sum := 1
	*/
	sum := 0
	for g := 1; g <= 100; g++ {
		sum += g
	}
	fmt.Println(sum)

	//3.打印1-100之间所有是9的倍数的整数的个数及总和
	var sum1 = 0
	var count = 0
	for d := 1; d <= 100; d++ {
		if d%9 == 0 {
			fmt.Println(d)
			sum1 += d
			count++
		}
	}
	fmt.Println(sum1)
	fmt.Println(count)

	//4.计算5的阶乘（12345 n 的阶乘 12.....n）1*2*3*4
	/*
	1.sum2=1*1
	2.sum2=sum2*2 1*1*2
	3.sum2=sum2*3 1*1*2*3
	 */
	sum2 := 1
	for s := 1; s <= 5; s++ {
		sum2 *= s
	}
	fmt.Println(sum2)

	//5.打印一个矩形(for循环的嵌套)
	/*
	****
	****
	****
	 */
	for e := 1; e <=12; e++ {
		fmt.Print("*")
		if e % 4 == 0{
			fmt.Println("")
		}
	}

	//6.for循环的嵌套
	/*
	for循环嵌套的一个执行流程
	w=0 打印4个* 一个换行
	w=1 打印4个* 一个换行
	w=2 打印4个* 一个换行
	w=3 跳出循环
	 */
	var row = 3
	var col = 4
	for w := 0; w < row; w++ {
		for v := 0; v < col; v++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	//7.打印一个三角形
	var li = 5
	for x := 1; x <= li; x++ {
		for y := 1; y <= x; y++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	//8.打印99乘法表
	for n := 1; n <= 9; n++ {
		for m := 1; m <= n; m++ {
			fmt.Printf("%v*%v=%v \t", m, n, n*m)
		}
		fmt.Println("")//打印空行
	}
}