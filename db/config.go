package db

import "strconv"

var RealmDBConf = &DbConfig{
	Host:     "localhost",
	Port:     3306,
	User:     "root",
	Pwd:      "mysql",
	Database: "mircore_realm",
}

var WorldDBConf = &DbConfig{
	Host:     "localhost",
	Port:     3306,
	User:     "root",
	Pwd:      "mysql",
	Database: "mircore_world",
}

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Pwd      string
	Database string
}

func (c *DbConfig) String() string {
	//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	return c.User + ":" + c.Pwd + "@tcp(" + c.Host + ":" + strconv.Itoa(c.Port) + ")/" + c.Database + "?charset=utf8&parseTime=True&loc=Local"
}
