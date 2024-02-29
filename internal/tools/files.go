package tools

import "os"

// IsFileExists 判断当前目录下是否存在名为fileName的文件
func IsFileExists(fileName string) bool {
	// 使用Stat获取文件的状态信息
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// 如果返回的错误类型为IsNotExist，则说明文件不存在
		return false
	}
	// 如果没有错误，说明文件存在
	return true
}
