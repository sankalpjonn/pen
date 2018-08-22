package pen

import (
	"testing"
	"time"
)

func TestPen(t *testing.T) {
	pen := New("TEST", "/tmp/pen-test-log.%Y%m%d-%S")

	pen.Write("pen-test-1")
	time.Sleep(time.Second * 2)

	pen.Write("pen-test-2")
	time.Sleep(time.Second * 2)

	pen.Lid()
}
