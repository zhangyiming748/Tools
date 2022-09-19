package download

import "testing"

func TestNewDownloader(t *testing.T) {
	var c Case
	c.resume = true
	c.concurrency = 4
	c.url = "https://f.bmcx.com/file/suijimimashengcheng/pic.jpg"
	c.output = "."
	if err := NewDownloader(c.concurrency, c.resume).Download(c.url, c.output); err != nil {
		t.Logf("error is %s\n", err.Error())
	}
}
