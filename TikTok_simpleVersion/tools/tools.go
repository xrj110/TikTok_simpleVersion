package tools

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	return hex.EncodeToString(hasher.Sum(nil))

}
