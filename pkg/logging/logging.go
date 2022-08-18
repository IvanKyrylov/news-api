package logging

import (
	"log"
	"os"
)

func Init() *log.Logger {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	return log.New(os.Stdout, "Default Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
