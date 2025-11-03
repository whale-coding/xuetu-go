package utils

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

/**
签名工具
*/

// CheckSignature 校验签名
func CheckSignature(sig, timestamp, nonce, token string) bool {
	arr := []string{token, timestamp, nonce}
	sort.Strings(arr)
	raw := strings.Join(arr, "")
	h := sha1.New()
	h.Write([]byte(raw))
	return fmt.Sprintf("%x", h.Sum(nil)) == sig
}
