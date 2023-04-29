package system

import (
	"os"
	"sync"
)

const ChainName = "GRIDC"

var once sync.Once
var pid int

func Getpid() int {
	once.Do(func() {
		pid = os.Getpid()
	})
	return pid
}
