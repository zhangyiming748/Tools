package amr

import (
	"github.com/zhangyiming748/amr2mp3"
	"testing"
)

//只用做演示
func TestShell(t *testing.T) {
	src := "/Users/zen/Github/Tools/amr/before"
	dst := "/Users/zen/Github/Tools/amr/after"
	amr2mp3.ConvertMP3(src, dst)
}
