package HASH

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func SHA256File(path string) (string, error) {
	h := sha256.New()
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
func SHA256(s ...string) {
	report := make(map[string]string)
	for _, v := range s {
		report[v], _ = SHA256File(v)
	}
	for k, v := range report {
		log.Printf("文件%v的SHA256值为%v\n", k, v)
	}
}
