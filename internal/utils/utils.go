package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

func UnixToWindowsPath(unixPath string) string {
	// Convert Unix-style path to Windows path
	windowsPath := filepath.FromSlash(unixPath)

	// If you're running on Windows, you might also want to convert
	// the path to an absolute path for better compatibility
	if os.PathSeparator == '\\' {
		return "C:\\Users\\cqt28\\Desktop\\snippetbox\\" + windowsPath
	}

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	return basepath + windowsPath
}
