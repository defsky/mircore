package log

import (
	"log"
	"os"
)

//RealmLog logger for realmd
var RealmLog *log.Logger

//WorldLog logger for world
var WorldLog *log.Logger

func init() {
	RealmLog = log.New(os.Stdout, "[Realm]", log.LstdFlags)
	WorldLog = log.New(os.Stdout, "[World]", log.LstdFlags)
}
