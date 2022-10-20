package helper

import "fmt"

func Int64SliceToStringSlice(val []int64) []string {
	strs := make([]string, 0)
	for _, v := range val {
		strs = append(strs, fmt.Sprintf("%v", v))
	}
	return strs
}
