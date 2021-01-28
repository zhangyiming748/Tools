package getFileType

import (

	"testing"
)

func TestMaster(t *testing.T) {
	var fp string = "../atom.mp4"
	p:=Master(fp)
	t.Log(p)
}