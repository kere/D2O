package model

import (
	"encoding/base64"

	"github.com/kere/gno/libs/util"
)

// Dasit 加密
func Dasit(str string) string {
	if str == "" {
		return ""
	}
	b64 := make([]byte, base64.StdEncoding.EncodedLen(len(str)))
	base64.StdEncoding.Encode(b64, util.Str2Bytes(str))

	l := len(b64)
	for i := 0; i < l; i++ {
		v := b64[i]
		b64[i] = byte(int(v) ^ ((i % 7 << 4) + (i % 15)))
	}

	return base64.StdEncoding.EncodeToString(b64)
}

// UnDasit 解密
func UnDasit(s string) []byte {
	b, _ := base64.StdEncoding.DecodeString(s)
	l := len(b)
	for i := 0; i < l; i++ {
		b[i] = byte(int(b[i]) ^ ((i % 7 << 4) + (i % 15)))
	}
	dst := make([]byte, base64.StdEncoding.DecodedLen(l))
	base64.StdEncoding.Decode(dst, b)
	l = len(dst)
	n := 0
	for i := l - 1; i > -1; i-- {
		if dst[i] != 0 {
			break
		}
		n++
	}
	if n == 0 {
		return dst
	}
	return dst[:l-n]
}
