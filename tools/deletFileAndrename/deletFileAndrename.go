package deletFileAndrename

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func checkIgnore(path string, ignores []string) bool {
	for _, ignore := range ignores {
		matched, _ := filepath.Match(ignore, filepath.Base(path))
		if matched {
			return true
		}
	}
	return false
}

func main() {
	root := "./kusion" // 指定的根目录
	ignoreFile := filepath.Join(root, ".gitignore")
	var ignores []string
	if _, err := os.Stat(ignoreFile); !os.IsNotExist(err) {
		file, err := os.Open(ignoreFile)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ignores = append(ignores, scanner.Text())
		}
	}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if checkIgnore(path, ignores) {
				return nil
			}
			// 如果是文件，则删除其内容
			err := ioutil.WriteFile(path, []byte{}, 0)
			if err != nil {
				return err
			}
			// 添加 .md 后缀
			newPath := path + ".md"
			err = os.Rename(path, newPath)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Success")
	}
}
