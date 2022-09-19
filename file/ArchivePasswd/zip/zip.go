package zip

import (
	"github.com/yeka/zip"
	"io/ioutil"
	"log"
	"os"
)

func UnZip(file, passwd string) {
	r, err := zip.OpenReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		if f.IsEncrypted() {
			f.SetPassword(passwd)
		}
		r, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		buf, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(f.Name, os.ModePerm)
		} else {
			ioutil.WriteFile(file+"/"+f.Name, buf, os.ModePerm)
		}
	}
}
