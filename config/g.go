package config

import (
	"log"
	"runtime"
)

// change log:
const (
	OFFICIAL_VERSION = "0.0.1"
	DADA_VERSION     = "1.2.0"
	VERSION          = DADA_VERSION
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
