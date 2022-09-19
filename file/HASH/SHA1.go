package HASH

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
)

//sha1
func SHA1File(path string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
func SHA1(s ...string) {
	report := make(map[string]string)
	for _, v := range s {
		report[v], _ = SHA1File(v)
	}
	for k, v := range report {
		log.Printf("文件%v的SHA1值为%v\n", k, v)
	}
}
