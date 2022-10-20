package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

// SplitFile 按照指定大小分割文件
// path：需要分割的文件路径
// size：按照多大来进行分割【这里写死4MB分割】
// return 分割后的文件的md5切片，分割后的文件路径，一个用来清理生成的临时目录回调函数，产生的错误
func SplitFile(path string) (md5s []string, tmps []string, clean func() error, err error) {
	const size = 1024 * 1024 * 4
	path, err = filepath.Abs(path)
	if err != nil {
		return nil, nil, nil, errors.WithStack(err)
	}
	tempDir, err := os.MkdirTemp(".", "")
	if err != nil {
		return nil, nil, nil, errors.WithStack(err)
	}
	info, err := os.Stat(path)
	if err != nil {
		return nil, nil, nil, errors.WithStack(err)
	}

	f, err := os.Open(path)
	if err != nil {
		if err != nil {
			return nil, nil, nil, errors.WithStack(err)
		}
	}
	defer f.Close()
	buffer := make([]byte, size)
	for i := int64(0); i <= info.Size()/size; i++ {

		n, err := f.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			return nil, nil, nil, errors.WithStack(err)
		}
		err = os.WriteFile(fmt.Sprintf("%s/%s_%d", tempDir, info.Name(), i+1), buffer[:n], os.ModePerm)
		if err != nil {
			return nil, nil, nil, errors.WithStack(err)
		}
		tmps = append(tmps, fmt.Sprintf("%s/%s_%d", tempDir, info.Name(), i+1))
		md5s = append(md5s, fmt.Sprintf("%x", md5.Sum(buffer[:n])))
	}

	return md5s, tmps, func() error {
		err := os.RemoveAll(tempDir)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	}, nil
}

// FileTime 获取文件的创建时间，最后访问时间，修改时间
func FileTime(info os.FileInfo) (string, string, string) {
	switch runtime.GOOS {
	case "windows":
		win := info.Sys().(*syscall.Win32FileAttributeData)
		return fmt.Sprint(win.CreationTime.Nanoseconds() / 1e9),
			fmt.Sprint(win.LastAccessTime.Nanoseconds() / 1e9),
			fmt.Sprint(win.LastWriteTime.Nanoseconds() / 1e9)
		//case "linux":
		//	linux := info.Sys().(*syscall.Stat_t)
		//	return fmt.Sprint(linux.Ctim.Nano() / 1e9),
		//		fmt.Sprint(linux.Atim.Nano() / 1e9),
		//		fmt.Sprint(linux.Mtim.Nano() / 1e9)
	}
	return "", "", ""
}

// PathSeparatorFormat 将所有的\替换为/
func PathSeparatorFormat(s string) string {
	return strings.Replace(s, "\\", "/", -1)
}
