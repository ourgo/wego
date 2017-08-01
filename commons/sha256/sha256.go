package sha256

import (
	"encoding/hex"
	"crypto/sha256"
)

// 加密密码
func ToHash(s string,k string) string {
	hanher := sha256.New()
	// 写入字符串
	hanher.Write([]byte(s))
	// 加盐
	hashed := hanher.Sum([]byte(k))
	// 返回16进位字符串
	return hex.EncodeToString(hashed)
}

// 密码对比
func ReturnBool(s ,k ,p string) bool {
	// 加密
	hashed := ToHash(s,k)
	// 返回
	if hashed==p {
		return true
	}else{
		return false
	}
}
