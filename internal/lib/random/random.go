package random

import (
	"math/rand"
	"strings"
	"time"
)

const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)
var src = rand.NewSource(time.Now().UnixNano())

func NewRandomString(n int) string {
  sb := strings.Builder{}
  sb.Grow(n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(symbols) {
      sb.WriteByte(symbols[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return sb.String()
}
