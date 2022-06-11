package utils

import (
	"math/rand"
	"time"
)

const uidLen = 10
const accessTokenLen = 20
const uidPrefix = "bu_"
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charsetNumber = "0123456789"

func genRandomString(prefix string, strLen int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, strLen)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return prefix + string(b)
}
func GenerateOAccessToken() string {
	return genRandomString("ot_", accessTokenLen, charset)
}
func GenerateUID() string {
	return genRandomString(uidPrefix, uidLen, charset)
}
func GenerateVerifyCode() string {
	return genRandomString("", 6, charsetNumber)
}
