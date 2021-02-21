package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {

	cmd := exec.Command("/usr/local/bin/youtube-dl","--proxy","127.0.0.1:8889","-F","youtube.com/watch?v=fX-TXWWf8rI")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error(), stderr.String())
	} else {
		log.Println(out.String())
	}

	var vcode string
	var acode string
	fmt.Println("which videocode?")
	fmt.Scanln(&vcode)
	fmt.Println("which audiocode?")
	fmt.Scanln(&acode)
	result:=strings.Join([]string{vcode,acode},"+")
	fmt.Println(result)

	var cmd2 *exec.Cmd
	var dl []byte
	var erro error
	cmd2 = exec.Command("/usr/local/bin/youtube-dl","--proxy","127.0.0.1:8889","-f",result,"youtube.com/watch?v=fX-TXWWf8rI")

	//cmd2 = exec.Command("whoami")
	// or cmd = exec.Command("bash", "-c", "put your commands here")
	dl, erro = cmd2.Output()

	if erro==nil{
		fmt.Println(string(dl))
	}else{
		fmt.Println(err)
	}



}
