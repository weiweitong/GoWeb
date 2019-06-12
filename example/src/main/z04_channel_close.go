package main

import "fmt"

func main()  {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 10; j++ {
		fmt.Println("sent job", j)
		jobs <- j
	}
	// 关闭一个通道意味着不再发送值了
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
