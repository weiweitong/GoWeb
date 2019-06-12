package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {

	dat, err := ioutil.ReadFile("./data/dat.txt")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open("./data/dat.txt")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	f.Close()

	fmt.Printf("\n\nWrite: \n")
	d1 := []byte("hello\ngo\n")
	err = ioutil.WriteFile("./data/dat.txt", d1, 0644)
	check(err)

	f, err = os.Create("./data/dat2.txt")
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err2 := f.Write(d2)
	check(err2)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err3 := f.WriteString("writes\n")
	check(err3)
	fmt.Printf("wrote %d bytes\n", n3)

	w := bufio.NewWriter(f)
	n4, err4 := w.WriteString("buffered\n")
	check(err4)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}
