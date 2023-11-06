package log

import (
	"log"
	"os"
)

var Error *log.Logger
var Info *log.Logger

func init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
}
