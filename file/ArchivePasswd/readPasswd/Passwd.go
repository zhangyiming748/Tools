package readPasswd

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadPasswd2Slice(passwd string) []string {
	passwds := []string{}
	fi, err := os.Open(passwd)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return []string{}
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		//fmt.Println(string(a))
		passwds = append(passwds, string(a))
	}
	return passwds
}
