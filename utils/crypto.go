package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// Verify ...
func Verify(t string, expire int64) []string {
	if len(t) == 0 {
		return nil
	}
	r := strings.Split(t, ".")
	timestamp, err := strconv.ParseInt(r[0], 10, 64)
	if err != nil {
		return nil
	}
	if len(r) < 2 || timestamp < time.Now().UnixNano()/int64(time.Millisecond)-expire {
		return nil
	}
	conf := GetConfig(r[1])

	if conf == "" {
		return nil
	}
	hash := hmac.New(sha256.New, []byte(conf))

	for i := 0; i < len(r)-1; i++ {
		hash.Write([]byte(r[i]))
	}
	if r[len(r)-1] != base64.StdEncoding.EncodeToString(hash.Sum(nil)) {
		return nil
	}
	return r
}

// HashFileMd5 ...
func HashFileMd5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}
