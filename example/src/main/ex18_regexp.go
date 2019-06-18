package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main()  {

	// 正则表达式，+号前面的字符至少出现一次
	plus_match0, _ := regexp.MatchString("runoo+b", "runob")
	plus_match1, _ := regexp.MatchString("runoo+b", "runoob")
	plus_match2, _ := regexp.MatchString("runoo+b", "runooob")
	fmt.Printf("+ match 不出现 %v 出现一次 %v 出现两次 %v\n\n", plus_match0, plus_match1, plus_match2)

	// * 号字符可不出现，或出现一次或者多次
	multi_match0, _ := regexp.MatchString("runoo*b", "runob")
	multi_match1, _ := regexp.MatchString("runoo*b", "runoob")
	multi_match2, _ := regexp.MatchString("runoo*b", "runooob")
	fmt.Printf("* match 不出现 %v 出现一次 %v 出现两次 %v\n\n", multi_match0, multi_match1, multi_match2)

	// ? 号字符可以匹配color或者colour, 问号前字符最多出现一次
	ques_match0, _ := regexp.MatchString("colou?r", "color")
	ques_match1, _ := regexp.MatchString("colou?r", "colour")
	ques_match2, _ := regexp.MatchString("colou?r", "colouur")
	fmt.Printf("? match 不出现 %v 出现一次 %v 出现两次 %v\n\n", ques_match0, ques_match1, ques_match2)

	// 测试一个字符串是否符合一个表达式
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Printf("字符串是否符合表达式 %v\n\n", match)

	// 上面直接使用字符串，但对于一些其他的正则任务
	// 需要Compile一个优化的Regexp的结构体
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Printf("匹配测试 %v\n\n", r.MatchString("peach"))

	fmt.Printf("查找匹配字符串 %v\n\n", r.FindString("peach punch"))

	// 查找匹配字符串，但是返回的匹配开始和结束位置索引而不是匹配的内容
	fmt.Printf("Find Index is %v\n\n", r.FindStringIndex("peach punch"))

	// 返回完全匹配和局部匹配的字符串
	fmt.Printf("Find String submatch %v\n\n", r.FindStringSubmatch("peach punch"))

	// 返回完全匹配和局部匹配的索引位置
	fmt.Printf("Find String Submatch Index %v", r.FindStringSubmatchIndex("peach punch"))

	// 带All的这个函数返回所有的匹配项，而不仅仅是首次匹配
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// All 同样可以对应到上面所有函数
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// 这里提供一个正整数来限制匹配次数
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 也可以使用[]byte 参数并将String 从函数中划掉
	fmt.Println(r.Match([]byte("peach")))


	// 创建正则表示式常量时，可以使用MustCompiler只返回一个值
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// regexp 包也可以用来替换部分字符串为其他值
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Func变量允许传递匹配内容到一个给定的函数中
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

	/*
	^[0-9]+abc$
	^为输入字符串的开始标记
	[0-9]+匹配多个数字，
	 */

}
