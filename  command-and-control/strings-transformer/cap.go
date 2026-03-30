package stringstransformer

import "strings"

func Cap(s string) string {
	return strings.Title(strings.ToLower(s))
}
