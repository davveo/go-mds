package conf

import "time"

type C struct {
	Mysql  Mysql
	Redis  Redis
	Server Server
}

type Mysql struct {
	DbName      string
	Host        string
	Port        string
	User        string
	Password    string
	MaxIdleConn int
	MaxOpenConn int
	ShowSQL     bool
}

type Redis struct {
	Host        string
	Port        int
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type Server struct {
	Port           int
	ReadTimeOut    time.Duration
	WriteTimeOut   time.Duration
	MaxHeaderBytes int
}
