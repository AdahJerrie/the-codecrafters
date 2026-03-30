package baseconverter

import (
	"fmt"
	"strconv"
)

func HexTodecimal(s string) int64 {
	result, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println(err)

	}
	return result

}
