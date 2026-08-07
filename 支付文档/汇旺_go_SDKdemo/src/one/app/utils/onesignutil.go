package utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"sort"
	"strings"
)

// Sign 用 RSA 私钥对参数对象进行签名
func Sign(data map[string]interface{}, privateKeyPEM string) (string, error) {
	signingStr := buildSignContent(data)
	log.Println("signingStr:", signingStr)
	// 解码私钥（你传进来的应该是 纯 Base64 编码）
	keyBytes, err := base64.StdEncoding.DecodeString(privateKeyPEM)
	if err != nil {
		return "", errors.New("私钥 Base64 解码失败: " + err.Error())
	}

	// 解析 PKCS#8 私钥
	privateKey, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		return "", errors.New("私钥解析失败: " + err.Error())
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("私钥格式错误")
	}

	hashed := sha256.Sum256([]byte(signingStr))
	signature, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// Verify 用 RSA 公钥验证签名是否正确
func Verify(signBase64 string, data map[string]interface{}, publicKeyPEM string) (bool, error) {
	signingStr := buildSignContent(data)
	log.Println("signingStr: ", signingStr)

	// 解码公钥（你传进来的是 Base64 编码的 PKIX 格式）
	keyBytes, err := base64.StdEncoding.DecodeString(publicKeyPEM)
	if err != nil {
		return false, errors.New("公钥 Base64 解码失败: " + err.Error())
	}

	// 解析 PKIX 公钥
	pubInterface, err := x509.ParsePKIXPublicKey(keyBytes)
	if err != nil {
		return false, errors.New("公钥解析失败: " + err.Error())
	}

	rsaPub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return false, errors.New("公钥格式错误")
	}

	// 解码签名
	signature, err := base64.StdEncoding.DecodeString(signBase64)
	if err != nil {
		return false, errors.New("签名 Base64 解码失败: " + err.Error())
	}

	// 哈希原始字符串
	hashed := sha256.Sum256([]byte(signingStr))

	// 验签
	err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return false, errors.New("验签失败: " + err.Error())
	}
	return true, nil
}

// 构建签名前的排序字符串，例如 key1=val1&key2=val2
func buildSignContent(data map[string]interface{}) string {
	keys := make([]string, 0, len(data))
	kvMap := make(map[string]string)

	for k, v := range data {
		if strings.EqualFold(k, "sign") || v == nil || strings.TrimSpace(toString(v)) == "" {
			continue
		}

		// 转换值为字符串
		var valStr string
		switch val := v.(type) {
		case []string:
			b, _ := json.Marshal(val)
			valStr = string(b)
		default:
			valStr = toString(v)
		}
		kvMap[k] = valStr
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var buf strings.Builder
	for i, k := range keys {
		buf.WriteString(k + "=" + kvMap[k])
		if i < len(keys)-1 {
			buf.WriteString("&")
		}
	}
	return buf.String()
}

func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return strings.TrimSpace(val)
	case float64:
		// Go 默认 JSON 解析数字为 float64，这里强转为高精度 big.Float，避免精度丢失
		bf := new(big.Float).SetFloat64(val)
		// 去除多余0，防止科学计数法
		return bf.Text('f', -1)
	case json.Number:
		// 如果你用 json.Decoder.UseNumber()，会遇到这个类型
		return val.String()
	default:
		return strings.TrimSpace(fmt.Sprintf("%v", val))
	}
}
