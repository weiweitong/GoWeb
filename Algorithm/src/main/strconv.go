package main

import (
	"fmt"
	"strconv"
)

func t(i interface{}) string {
	switch i.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case float64:
		return "float64"
	case uint:
		return "uint"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func main()  {
	s1 := make([]byte, 0, 100)
	s1 = strconv.AppendInt(s1, 4567, 10)
	fmt.Println(string(s1))

	s1 = strconv.AppendBool(s1, false)
	fmt.Println(string(s1))
	fmt.Println(s1)

	s1 = strconv.AppendQuote(s1, "abcdefg")
	s1 = strconv.AppendQuoteRune(s1, '单')
	fmt.Println(string(s1))

	fmt.Println("\n")
	a2 := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a2, b, c, d, e)
	fmt.Println(t(e))

	ax, _ := strconv.ParseBool("false")
	fmt.Println(t(ax))

}
