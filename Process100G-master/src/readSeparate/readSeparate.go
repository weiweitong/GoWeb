package readSeparate

import (
	"bufio"
	"io"
	"os"
	"randHash"
	"strconv"
)

// 读取文件到内存中，到了SizeBatch就使用handle函数处理memString
// 之后刷新memString避免内存爆炸
func ReadFile(filePath string, handle func([]string, string, int), BatchSize, NumFile int, tmpPath string) error {
	f, err := os.Open(filePath)
	defer f.Close()
	checkError(err)
	buf := bufio.NewReader(f)

	// 将小于等于一个BatchSize的url保存到memString中
	memString := make([]string, 0)
	for count := 0; ; count++ {
		line, _, err := buf.ReadLine()
		if count == BatchSize {
			handle(memString, tmpPath, NumFile)
			memString = make([]string, 0)
			count = 0
		}
		memString = append(memString, string(line))
		if err != nil {
			if err == io.EOF {
				if len(memString) > 0 {
					handle(memString, tmpPath, NumFile)
					memString = make([]string, 0)
					count = 0
				}
				return nil
			}
			return err
		}
	}
}

// 使用Hash方程分割url切片然后把它们放入NumFile个文件中
func SetPartition(strs []string, tmpPath string, NumFile int) {
	fileMap := make(map[string][]string)
	for _, str := range strs {
		if str == "" {
			continue
		}
		partition := tmpPath + "/" + strconv.Itoa(int(randHash.TongHash64(str)) % NumFile) + ".txt"
		if _, ok := fileMap[partition]; ok {
			fileMap[partition] = append(fileMap[partition], str)
		} else {
			fileMap[partition] = []string{str}
		}
	}
	_, err := os.Stat(tmpPath)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(tmpPath, os.ModePerm)
			checkError(err)
		}
	}
	for k, vs := range fileMap {
		f, err := os.OpenFile(k, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		checkError(err)
		for _, v := range vs {
			_, err = f.Write([]byte(v + "\n"))
		}
		f.Close()
	}
}


func checkError(err error) {
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
}


