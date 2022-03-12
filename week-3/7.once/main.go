//package main
//
//import "sync"

//type DbConnection struct{}
//
//var (
//	conn *DbConnection
//)
//
//func GetConnection() *DbConnection {
//	if conn == nil {
//		conn = &DbConnection{}
//	}
//	return conn
//}
//

package main

import (
	"sync"
)

type DbConnection struct{}

var (
	dbConnOnce sync.Once
	conn       *DbConnection
)

func GetConnection() *DbConnection {
	dbConnOnce.Do(func() {
		conn = &DbConnection{}
	})
	return conn
}
