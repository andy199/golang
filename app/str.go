package main

import (
	"fmt"
	"strings"
)
//字符串类型

func main()  {
	first := "Ai"
	last := "Andy"
	name := first +" "+ last
	fmt.Println("my name is", name)

	//字符串转义符
	str2 := "c:\\go\\bin" //c:\go\bin
	fmt.Println(str2)

	//多行字符串 (反引号)tab键上面
	str1 := `this is str
	ok in
	and or
	sas df
	`
	fmt.Println(str1)

	//len(str)求长度
	var str3 = "你好"
	fmt.Println(len(str3))
	
	// + 或者 fmt.Sprintf拼接字符串
	str4 := "你好"
	str5 := "golang"
	str6 := str4 + str5
	fmt.Println(str6)
	
	str7 := "你好"
	str8 := "golang"
	str9 := fmt.Sprintf("%v %v", str7, str8)
	fmt.Println(str9)
	
	//strings.Split 分割字符串 strings需要引入strings包
	var str10  = "123-456-789"
	arr := strings.Split(str10, "-")//切片
	fmt.Println(arr)
	//join操作 表示把切片链接成字符串
	str11 := strings.Join(arr, "*")//连接
	fmt.Println(str11)

	//strings.contains判断是否包含，包含（true）不包含（false）
	str12 := "this is str"
	str13 := "this"
	flag := strings.Contains(str12, str13)
	fmt.Println(flag)//true


	//strings.HasSuffix, strings.HasSuffix 字符串前缀/后缀判断
	str14 := "this is str"
	str15 := "this"
	flag1 := strings.HasPrefix(str14, str15)
	fmt.Println(flag1)//true

	str16 := "this is str"
	str17 := "str"
	flag2 := strings.HasSuffix(str16, str17)
	fmt.Println(flag2)//true


	//strings.Index(), strings.LastIndex() 子串出现的位置，查找不到返回-1
	str18 := "this is str"
	str19 := "str"
	flag3 := strings.Index(str18, str19)
	fmt.Println(flag3)//8位置

	str20 := "this is str"
	str21 := "tr"
	flag4 := strings.LastIndex(str20, str21)
	fmt.Println(flag4)//9位置

}

