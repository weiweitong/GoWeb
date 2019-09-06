package verify

import (
	"bufio"
	"fmt"
	"os"
	"randHash"
	"regexp"
	"strconv"
)

func VerifyOutput(NumFile int)  {

	fmt.Printf("\n\n")

	f, _ := os.Open("./src/data/output.txt")
	buf := bufio.NewReader(f)

	for i := 0; i < NumFile; i++ {
		line, _, _ := buf.ReadLine()

		s := string(line)
		r1 := regexp.MustCompile("[0-9]+")
		r2:= regexp.MustCompile(`https:(.+)`)

		num := r1.FindString(s)
		frequency, _ := strconv.Atoi(num)
		//fmt.Println(frequency)
		url := r2.FindString(s)
		//fmt.Println(url)

		partition := "./src/tmp/" + strconv.Itoa(int(randHash.TongHash64(url)) % NumFile) + ".txt"
		//fmt.Println(partition)

		f2, _ := os.Open(partition)
		buf2 := bufio.NewReader(f2)
		count := 0
		for line2, _, err := buf2.ReadLine(); err == nil; line2, _, err = buf2.ReadLine() {
			if string(line2) == url {
				count++
			}
		}
		if count == frequency {
			fmt.Printf("Output %d is Right\n\n", i)
		} else {
			fmt.Printf("Output %d is Wrong\n\n", i)
		}
	}
}


