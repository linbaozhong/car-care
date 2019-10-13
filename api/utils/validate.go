package utils

import (
	"regexp"
	"strconv"
	"unicode/utf8"
)

var Validate = new(validator)

type validator struct{}

// Len 获取字符串长度
func (v *validator) Len(s string) int {
	return utf8.RuneCountInString(s)
}

// Integer 是否数字，并返回数字
func (v *validator) Integer(s string) (bool, int64) {
	if ok, _ := regexp.MatchString("^[0-9]+$", s); ok {
		val, _ := strconv.ParseInt(s, 10, 64)
		return true, val
	}
	return false, 0
}

// Range 范围
func (v *validator) Range(s, min, max int) bool {
	return s > min && s < max
}

// Mobile 手机号
func (v *validator) Mobile(s string) bool {
	ok, _ := regexp.MatchString(`^1[34578]\d{9}$`, s)
	return ok
}

// Email 邮箱
func (v *validator) Email(s string) bool {
	ok, _ := regexp.MatchString(`^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`, s)
	return ok
}
// Alphanumeric 字母数字
func (v *validator) Alphanumeric(s string) bool {
	ok, _ := regexp.MatchString(`^[\w]+$`, s)
	return ok
}
// Or 是否其中之一
func (v *validator) Or(s interface{}, args ...interface{}) bool {
	for _, val := range args {
		//switch t:=val.(type) {
		//case string:
		//case int:
		//}
		if val == s {
			return true
		}
	}
	return false
}
