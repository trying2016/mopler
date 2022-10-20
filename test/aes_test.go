package test

import (
	"crypto/aes"
	"github.com/766800551/mopler/helper"
	"testing"
)

func TestEncrypt(t *testing.T) {
	err := helper.AesEncrypt("C:\\Users\\76680\\Desktop\\新建文本文档.txt",
		[]byte("2pfkl36zyctsc0ki"),
		[aes.BlockSize]byte{'c', 's', 'd', 'r', 'g', 'e', 'x', 'y', 's', 'h', 'l', '6', 'u', 'b', '6', 'y'})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecrypt(t *testing.T) {
	err := helper.AesDecrypt("C:\\Users\\76680\\GolandProjects\\mopler\\test\\eee\\[aes_ofb]chromedriver_win32.zip",
		[]byte("2pfkl36zyctsc0ki"),
		[aes.BlockSize]byte{'c', 's', 'd', 'r', 'g', 'e', 'x', 'y', 's', 'h', 'l', '6', 'u', 'b', '6', 'y'})
	if err != nil {
		t.Fatal(err)
	}
}
