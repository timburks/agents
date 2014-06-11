package agents

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

func MD5HashWithSalt(input, salt string) string {
	hasher := hmac.New(md5.New, []byte(salt))
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

