package theme

import (
	"os"
	"path/filepath"
)

func IntegrityCheck(rootDir string) {
	ensureBaseJetFiles(rootDir, false)
}

func ensureBaseJetFiles(dir string, hasParent bool) {
	// 检查当前目录是否有base.jet
	currentBaseJetPath := filepath.Join(dir, "base.jet")
	_, err := os.Stat(currentBaseJetPath)

	if os.IsNotExist(err) {
		// 如果当前目录没有base.jet
		content := ""
		if hasParent {
			// 如果这个目录不是根目录，则新的base.jet继承自上一级目录的base.jet
			content = `{{ extends "../base.jet" }}`
		}
		// 创建base.jet文件
		os.WriteFile(currentBaseJetPath, []byte(content), 0775)
	}

	// 遍历当前目录
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			// 递归检查每个子目录，子目录有父目录，所以传入true
			ensureBaseJetFiles(filepath.Join(dir, entry.Name()), true)
		}
	}
}
