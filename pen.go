package pen

import (
	"github.com/lestrrat-go/strftime"
	"log"
	"os"
	"time"
)

type Pen struct {
	// thread save logger instance
	logger *log.Logger

	// file
	fp *os.File

	// log prefix
	prefix string

	// location of file
	location *strftime.Strftime

	// time location
	tloc *time.Location

	// file name created by builder
	path string
}

func New(prefix string, location string) *Pen {
	p, err := strftime.New(location)
	if err != nil {
		panic(err)
	}

	return &Pen{
		location: p,
		tloc:     time.Local,
		prefix:   prefix,
	}
}

func (self *Pen) initLogger() error {
	ct := time.Now()
	path := self.location.FormatString(ct.In(self.tloc))

	if self.logger == nil {
		fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		self.fp = fp
		self.logger = log.New(self.fp, "", 0)

		self.path = path
	} else if self.path != path {
		if self.fp != nil {
			self.fp.Close()
		}
		self.fp = nil
		self.logger = nil
		return self.initLogger()
	}
	return nil
}

func (self *Pen) Write(content string) error {
	err := self.initLogger()
	if err != nil {
		return err
	}
	self.logger.Println(time.Now().Format("2006-01-02 15:04:05")+" ["+self.prefix+"] ", content)
	return nil
}

func (self *Pen) Lid() {
	if self.fp != nil {
		self.fp.Close()
		self.logger = nil
	}
}
