package threads

import (
	"runtime"
	"strings"
	"testing"
)

func TestThreads(t *testing.T) {
	sysType := runtime.GOOS
	arch := runtime.GOARCH
	if sysType == "linux" {
		t.Logf("LINUX系统\n")
		if strings.Contains(arch, "arm") {
			t.Logf("arm64架构\n")
		} else {
			t.Logf("amd64架构\n")
		}
	}
	if sysType == "windows" {
		t.Logf("windows系统\n")
		if strings.Contains(arch, "arm") {
			t.Logf("arm64架构\n")
		} else {
			t.Logf("amd64架构\n")
		}
	}
	if sysType == "darwin" {
		t.Logf("mac系统\n")
		if strings.Contains(arch, "arm") {
			// 在MacBookAir上正常显示
			t.Logf("arm64架构\n")
		} else {
			// 在MacBookPro上正常显示
			t.Logf("amd64架构\n")
		}
	}
}
