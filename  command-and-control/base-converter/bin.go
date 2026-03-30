package baseconverter

import (
	"fmt"
	"strconv"
)

func BinTodecimal(s string) int64 {
	result, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return result
}
