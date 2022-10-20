package test

import (
	"fmt"
	"github.com/766800551/mopler/helper"
	"os"
	"testing"
)

func TestSplitFile(t *testing.T) {
	md5s, _, clean, err := helper.SplitFile("C:\\Users\\76680\\GolandProjects\\mopler\\chromedriver_win32.zip")
	if err != nil {
		t.Fatal(err)
	}
	err = clean()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%q", md5s)
}

func TestFileTime(t *testing.T) {
	f, _ := os.Stat("C:\\Users\\76680\\GolandProjects\\mopler\\chromedriver_win32.zip")
	fmt.Println(helper.FileTime(f))
}
