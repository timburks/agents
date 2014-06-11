package main

import (
    "crypto/md5"
	"crypto/hmac"
    "encoding/hex"
	"fmt"
)

func md5HashWithSalt(input, salt string) string {
	hasher := hmac.New(md5.New, []byte(salt))
    hasher.Write([]byte(input))
    return hex.EncodeToString(hasher.Sum(nil))
}
 

func main () {
	salt := "agent.io"
	password := "123"
		
	s := md5HashWithSalt(password, salt)	
	fmt.Println(s)
		
}