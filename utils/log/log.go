package log

import (
	"log"
	"os"
)

var (
	//Core logger for core log
	Core *log.Logger

	//Realm logger for realmd log
	Realm *log.Logger

	//World logger for world log
	World *log.Logger
)

func init() {
	Core = log.New(os.Stdout, "[Core ]", log.LstdFlags)
	Realm = log.New(os.Stdout, "[Realm]", log.LstdFlags)
	World = log.New(os.Stdout, "[World]", log.LstdFlags)
}
