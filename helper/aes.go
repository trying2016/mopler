// Package helper 通用的帮助方法
package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/pkg/errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// AesEncrypt 加密单个文件，加密后文件会多一个[aes_ofb]的前缀
func AesEncrypt(path string, key []byte, iv [aes.BlockSize]byte) error {
	prefix := "[aes_ofb]"
	inFile, err := os.Open(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer inFile.Close()
	block, err := aes.NewCipher(key)
	if err != nil {
		return errors.WithStack(err)
	}
	path, err = filepath.Abs(path)
	if err != nil {
		return errors.WithStack(err)
	}
	filename := filepath.Dir(path) + string(filepath.Separator) + prefix + filepath.Base(path)
	stream := cipher.NewOFB(block, iv[:])
	outFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return errors.WithStack(err)
	}
	defer outFile.Close()
	writer := &cipher.StreamWriter{S: stream, W: outFile}
	if _, err := io.Copy(writer, inFile); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// AesDecrypt 解密单个文件，解密前请确保文件有[aes_ofb]前缀
func AesDecrypt(path string, key []byte, iv [aes.BlockSize]byte) error {
	prefix := "[aes_ofb]"
	inFile, err := os.Open(path)
	if err != nil {
		return errors.WithStack(err)
	}
	if strings.Index(path, prefix) < 0 {
		return errors.New("该文件没有加密前缀")
	}
	defer os.Remove(path)
	defer inFile.Close()
	block, err := aes.NewCipher(key)
	if err != nil {
		return errors.WithStack(err)
	}
	stream := cipher.NewOFB(block, iv[:])
	outFile, err := os.OpenFile(strings.Replace(path, prefix, "", 1), os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return errors.WithStack(err)
	}
	defer outFile.Close()
	reader := &cipher.StreamReader{S: stream, R: inFile}
	if _, err := io.Copy(outFile, reader); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// AesDecryptRecursion 递归解密文件
func AesDecryptRecursion(path string, key []byte, iv [aes.BlockSize]byte) error {
	path, _ = filepath.Abs(path)
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		//如果不是目录，那么就解密
		if !info.IsDir() {
			err := AesDecrypt(path, key, iv)
			if err != nil {
				return err
			}
		}
		return err
	})
	return errors.WithStack(err)
}
