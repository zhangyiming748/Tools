package HASH

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func MD5File(path string) (string, error) {
	md5h := md5.New()
	m, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer m.Close()
	_, err = io.Copy(md5h, m)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5h.Sum(nil)), nil
}
func MD5(s ...string) {
	report := make(map[string]string)
	for _, v := range s {
		report[v], _ = MD5File(v)
	}
	for k, v := range report {
		log.Printf("文件%v的MD5值为%v\n", k, v)
	}
}
