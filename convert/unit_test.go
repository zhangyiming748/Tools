package convert

import (
	"fmt"
	"runtime"
	"testing"
)

func TestThreads(t *testing.T) {
	fmt.Printf("%d\n", runtime.NumCPU())
	fmt.Printf("%s\n", runtime.GOARCH)
	fmt.Printf("%s\n", runtime.GOOS)
}
