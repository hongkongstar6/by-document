package utils

import (
	"github.com/google/uuid"
	"strings"
)

// GetNonce 生成一个不带连字符的 UUID 字符串
func GetNonce() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
