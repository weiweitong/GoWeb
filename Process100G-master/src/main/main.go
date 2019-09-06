package main

import (
	"fileToHeap"
	"fmt"
	"genURL"
	"heapToFile"
	"readSeparate"
	"strconv"
	"time"
	"verify"
)

/*
100GB url 文件，使用 1GB 内存计算出出现次数 top100 的 url 和出现的次数

思路：
先将100G文件分为100个以上(设为N>100）的小于1G的小文件，将每个小文件加载到内存中，
用HashMap的形式进行映射，排序的话就用可以堆的形式吧，这样时间复杂度也低，
每个小文件建立一个堆，选出一百个次数最多的url，然后这N个小文件的TOP100一起读入内存，

原计划：
每次读入小文件堆的最大值(顶端)到有100个底层节点的归并树，俩俩比较，树的顶点就是最大值，
然后这个最大值的堆读入一个新的顶点，并读入树中，
这棵树只需要更新这个新值到root的路径上的节点即可，
这样重复100次，得到的就是100个最大的url

原计划弱点：每次更新树都要做一次对磁盘IO对一个文件堆进行读取
磁盘IO远比在内存中进行计算耗时

改进：
新建一个空的主堆，在每次从内存中读取小文件并建堆时，
将主堆与小文件生成的最小堆合并，得到100个出现次数最多的URL，
这样当主堆与所有小文件堆都合并后，最后的主堆中必定是出现次数最多的前100个URL
 */

func main() {

	// 将100G的大文件分为NumFile个小文件数
	var NumFile int = 100

	// 出现次数最多的NumTop个url
	var NumTop int = 100

	// IE浏览器最多支持2KB长度的url，考虑到兼容性不妨设定url最大长度为2k
	// 那么BatchSize * 2KB 应该小于1GB内存, 实际上url平均大小 ～ 200B，故不妨设置成Batchsize为3M
	var BatchSize int = 3 * 1024 * 1024

	filePath := "./src/data/dataset.txt"
	genURL.GenerateURL(filePath, 1024, 1024 * 1024)

	t2 := time.Now() // get current time

	tmpPath := "./src/tmp"
	err := readSeparate.ReadFile(filePath, readSeparate.SetPartition, BatchSize, NumFile, tmpPath)
	CheckError(err)

	heap := fileToHeap.Reduce(BatchSize, NumFile, NumTop, tmpPath)

	heapToFile.HeapToFile(heap)

	fmt.Println("The top " + strconv.Itoa(NumTop) + " results have been output to file \"./output.txt\"")
	elapsed := time.Since(t2)
	fmt.Println("App elapsed: ", elapsed)

	verify.VerifyOutput(NumFile)
}

func CheckError(err error) {
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
}





