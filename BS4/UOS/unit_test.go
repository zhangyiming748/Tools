package UOS

import (
	"testing"
)

func TestUOS(t *testing.T) {
	var urls = []string{"https://cdimage-download.chinauos.com",
		"https://cdimage-download.chinauos.com/education",
		"https://cdimage-download.chinauos.com/home",
		"https://cdimage-download.chinauos.com/home-ditch",
		"https://cdimage-download.chinauos.com/professional-wayland",
		"https://cdimage-download.chinauos.com/sp1-fix"}
	for _, url := range urls {
		UOS(url)
	}
}
