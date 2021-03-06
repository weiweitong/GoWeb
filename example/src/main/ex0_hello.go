// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.2.md
package main

import (
	"errors"
	"fmt"
)

func abd() int {
	return 1
}

func main() {
	var cmp complex64 = 5 + 5i
	fmt.Printf("Value is: %v", cmp)

	var s string = "hello"
	ch := []byte(s)
	ch[0] = 'c'
	s2 := string(ch)
	fmt.Printf("\nValue is: %v\n", s2)

	s = "hello"
	s = "c" + s[1:]

	m := `hello
				world`
	fmt.Printf("\nValue is: %v\n", m)

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}

	const i = 10
	var arr [i]int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Printf("The last element is %d\n", arr[9])

	a1 := [3]int{1, 2, 3}
	a2 := [10]int{1, 2, 3}
	a3 := [...]int{4, 5, 6}
	fmt.Printf("%v + %v + %v\n\n", a1, a2, a3)

	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Printf("%v + %v\n\n", doubleArray, easyArray)

	var fslice []int
	fmt.Printf("%v\n\n", fslice)

	slice := []byte{'a', 'b', 'c', 'd'}
	fmt.Printf("%c\n\n", slice)

	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var aSlice, bSlice []byte

	aSlice = array[:3]
	aSlice = array[5:]
	aSlice = array[:]

	aSlice = array[3:7]
	bSlice = aSlice[1:3]
	bSlice = aSlice[:3]
	bSlice = aSlice[0:5]
	bSlice = aSlice[:]

	fmt.Printf("%c + %c\n\n", aSlice, bSlice)

	// slice 是引用类型，如果改变其中元素的值时，其他所有引用都会改变该值
	// 如果修改了aSlice中元素的值，bSlice也会改变
	aSlice[1] = 'z'
	fmt.Printf("%c\n\n", bSlice)

	fmt.Println(abd())
	// 声明一个key是字符串，值为int的字典，这种方式的声明需要在使用前使用make初始化
	//var numbers map[string]int
	// 另一种map的声明方式
	//numbers
}
