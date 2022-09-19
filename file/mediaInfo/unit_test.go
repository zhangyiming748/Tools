package mediaInfo

import "testing"

func TestMediaInfo(t *testing.T) {
	src := "/Users/zen/Movies"
	file := "传统魔术吞剑.mp4"
	MediaInfo(src, file)
}
