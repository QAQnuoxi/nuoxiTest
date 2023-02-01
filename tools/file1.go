package tools

import "os"

func File1() {
	path := "all/directory"
	os.MkdirAll(path, 0777)
	userFile := path + "/astaxie.txt"
	files, _ := os.Create(userFile)
	files.Write([]byte("hgajkgdasgga"))
	files.WriteString("11111jhkadflbsgasdf")

	fl, _ := os.Open(userFile)
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		//os.Stdout.Write(buf[:n])
	}
}
