package test

import (
	"fmt"
	"testing"
)

// 获取用户信息
func TestUserInfo(t *testing.T) {
	u, err := sdk.UserInfo()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", u)
}
