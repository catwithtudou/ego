package helper

import (
	"encoding/hex"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/ebar-go/ego/http/constant"
	"hash/crc32"
	"encoding/base64"
)

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// GetHash 获取hash加密内容
func GetHashString(s string) string {
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(s))
	result := Sha1Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", result)
}

//UniqueId 生成Guid字串
func UniqueId() string {
	return uuid.NewV4().String()
}

// NewTraceId 生成全局ID
func NewTraceId() string {
	return constant.TraceIdPrefix + UniqueId()
}

// NewRequestId 生成请求ID
func NewRequestId() string {
	return constant.RequestIdPrefix + UniqueId()
}

func HashCode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func DecodeBase64Str(encoded string) string {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return ""
	}
	return string(decoded)
}
