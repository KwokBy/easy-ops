package file

import (
	"errors"
	"fmt"
	"os"
)

// IsExist checks if a file exists
func IsExist(fileAddr string) bool {
	_, err := os.Stat(fileAddr)
	return err == nil
}

// GetFilePath 向外寻找文件路径
func GetFilePath(fileAddr string, searchTime int) (string, error) {
	if IsExist(fileAddr) {
		return fileAddr, nil
	}
	if searchTime < 0 {
		return "", fmt.Errorf("file not exisit: %s", fileAddr)
	}
	return GetFilePath(fmt.Sprintf("../%s", fileAddr), searchTime-1)
}

// PathExists 判断文件目录是否存在
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
