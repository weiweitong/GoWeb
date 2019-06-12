package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct{
	x, y int
}


var pl = fmt.Println
var pf = fmt.Printf

func main()  {

	pl("Contains:		", strings.Contains("test", "es"))
	pl("Count:		", strings.Count("test", "t"))
	pl("HasPrefix:	", strings.HasPrefix("test", "te"))
	pl("HasSuffix:	", strings.HasSuffix("test", "st"))
	pl("Index:		", strings.Index("test", "e"))
	pl("Join:		", strings.Join([]string{"a", "b"}, "-"))
	pl("Repeat:		", strings.Repeat("a", 5))
	pl("Replace:	", strings.Replace("oofoo", "o", "0", -1))
	pl("Replace:	", strings.Replace("oofoo", "o", "0", 1))
	pl("Split:		", strings.Split("a-b-c-d-e", "-"))
	pl("ToLower:	", strings.ToLower("TEST"))
	pl("ToUpper:	", strings.ToUpper("test"))
	pl()

	pl("Len:			", len("hello"))
	pl("Char:			", "hello"[1])

	pf("\n\n")
	p := point{1, 2}
	pf("%v\n", p)

	pf("%+v\n", p)

	pf("%#v\n", p)

	pf("%T\n", p)

	pf("%t\n", true)

	pf("%d\n", 123)

	pf("%b\n", 14)

	pf("%c\n", 33)

	pf("%x\n", 456)

	pf("%f\n", 78.9)

	pf("%e\n", 1234000000.0)
	pf("%E\n", 1234000000.0)

	pf("%s\n", "\"string\"")

	pf("%q\n", "\"string\"")

	pf("%x\n", "hex this")

	pf("%p\n", &p)

	pf("|%6d|%6d|\n", 12, 345)

	pf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	pl(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")


}
