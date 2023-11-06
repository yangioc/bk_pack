package util

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/rs/xid"
)

// 生成亂數字串
func UnitStr() string {
	return xid.New().String()
}

// 移除無法顯示的 unicode
func RemoveUnPrintUncode(source string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, source)
}

// 從字串中取得數字
func GetNumberInString(source string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(source, -1)
}
