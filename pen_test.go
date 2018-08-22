package pen

import (
  "time"
  "testing"
)

func TestPen(t *testing.T) {
  pen := New("TEST", "/tmp/pen-test-log.%Y%m%d")

  pen.Write("pen-test-1")
  time.Sleep(time.Second * 2)

  pen.Write("pen-test-2")
  time.Sleep(time.Second * 2)

  pen.Lid()
}
