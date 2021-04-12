package tools

import (
	"bytes"
	"strings"
)

/*将字符转换为camel风格*/
func Camel(strRaw string) string {
	strList := strings.Split(strRaw, "_")
	var charBuffer bytes.Buffer
	for _, str := range strList {
		str = strings.Title(str)
		charBuffer.WriteString(str)
	}
	return charBuffer.String()
}
