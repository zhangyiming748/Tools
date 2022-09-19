package file

import (
	"testing"
)

func TestFilter(t *testing.T) {
	src := "exam.html"
	Filter(src, "jpg")
}
func TestGetName(t *testing.T) {
	long := "/Users/zen/Downloads/abc.iso"
	short := "abc.iso"
	longwithdot := "/Users/zen/Downloads/192.168.1.10.abc.iso"
	shortwithdot := "192.168.1.10.abc.iso"

	ret1 := LongNameGetFileName(long)
	ret2 := LongNameGetExtName(long)

	ret3 := LongNameGetFileName(longwithdot)
	ret4 := LongNameGetExtName(longwithdot)

	ret5 := LongNameGetFileName(short)
	ret6 := LongNameGetExtName(short)

	ret7 := LongNameGetFileName(shortwithdot)
	ret8 := LongNameGetExtName(shortwithdot)

	t.Logf("1=%v\n2=%v\n3=%v\n4=%v\n5=%v\n6=%v\n7=%v\n8=%v\n", ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8)

}
