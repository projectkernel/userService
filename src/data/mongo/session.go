package mongo

import (
	"gopkg.in/mgo.v2"
	"net"
	"crypto/tls"
	"time"
)

func Get(addr []string) (*mgo.Session, error){
	return GetWithCredentials(addr, "", "", "", false)
}

func GetWithCredentials(addr []string, user string, pass string, db string, ssl bool) (*mgo.Session, error) {
	dialInfo :=  &mgo.DialInfo{
		Addrs: addr,
		Database: db,
		Username: user,
		Password: pass,
		Timeout: 5 * time.Second,
	}
	if ssl {
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), &tls.Config{})
			return conn, err
		}
	}
	return mgo.DialWithInfo(dialInfo)
}