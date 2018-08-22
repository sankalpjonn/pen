# pen
Simple time based and thread safe logger. internally uses the log package

Inspired by
--------
https://github.com/arriqaaq/qalam


Features
--------
- Create a simple log file handler, will do log rotation automatically.
- Completely thread safe and can be used inside multiple go routines

Installing
----------

```
go get -u github.com/sankalpjonn/pen
```

Example
-------

The example below logs to a file with (%Y%m%d) format:

```go
package main

import (
	"github.com/sankalpjonn/pen"
)

func main() {
	pen := New("LOG-PREFIX", "/tmp/log.%Y%m%d")
	pen.Write("foo")
	pen.Write("bar")

  // to close the pen
	pen.Lid()
}

```

CONTACT
-------
[sankalpjonna@gmail.com](sankalpjonna@gmail.com)
