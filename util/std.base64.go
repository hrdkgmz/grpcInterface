package util

import "encoding/base64"

// EncodeBytes 把一个[]byte通过base64编码成string
func EncodeBase64Bytes(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// DecodeBytes 把一个string通过base64解码成[]byte
func DecodeBase64Bytes(src string) (s []byte, err error) {
	s, err = base64.StdEncoding.DecodeString(src)
	return
}

// EncodeBase64 把一个string通过base64编码
func EncodeBase64(src string) string {
	return EncodeBase64Bytes([]byte(src))
}

// Decode 把一个string通过base64解码
func DecodeBase64(src string) (s string, err error) {
	buf, err := DecodeBase64Bytes(src)
	if err != nil {
		return
	}
	s = string(buf)
	return
}
