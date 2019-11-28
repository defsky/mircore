package db

import "mircore/utils/log"

func Migrate() {
	log.Core.Println("DB migration start...")

	RealmDB.AutoMigrate(&Account{})

	log.Core.Println("DB migration complete")
}
