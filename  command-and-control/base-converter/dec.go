package baseconverter

import "fmt"

func Dec(s int64) string {
	result := fmt.Sprintf("Binary: %b\n Hex: %X\n", s, s)
	return result
}
