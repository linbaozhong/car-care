package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func _md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func PwdString(pwd string,salt string) string {
	return _md5(pwd + salt)
}

// 从请求tk中解析出用户ID和Token
func ParseIdTokenByTk(tk string) (int64, string) {
	pos := strings.Index(tk, "_")
	if pos < 1 {
		return 0, tk
	}
	token := tk[:pos]
	id, err := strconv.ParseInt(tk[pos+1:], 16, 64)
	if err == nil {
		return id, token
	}
	return 0, tk
}
