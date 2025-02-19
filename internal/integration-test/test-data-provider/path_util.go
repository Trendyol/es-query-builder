package testdataprovider

import (
	"path/filepath"
	"runtime"
)

func ResolvePath(relativePath string) string {
	_, currentFile, _, _ := runtime.Caller(0)
	dir := filepath.Dir(currentFile)
	return filepath.Join(dir, relativePath)
}
