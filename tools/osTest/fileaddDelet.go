package osTest

import (
	"os"
	"strconv"
)

/*
执行sum次
创建目录
*/
func MkdirDirectory(sum int) {
	//name := "test file" + strconv.Itoa(sum)
	for i := 0; i < sum; i++ {
		path := "all/directory" + strconv.Itoa(i)
		//os.Mkdir(name, 0777)
		os.MkdirAll(path, 0777)
	}
}
func CreatFileAll(sum int) {
	for i := 0; i < sum; i++ {
		path := "./all/directory" + strconv.Itoa(i) + "/md.txt"
		os.Create(path)
	}
}
func OpenFileAll(str string) {
	os.Open(str)
}
