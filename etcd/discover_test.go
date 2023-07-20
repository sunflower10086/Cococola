package etcd

import (
	"log"
	"testing"
)

func TestDiscover(t *testing.T) {
	var endpoints = []string{"localhost:2379"}
	ser := NewServiceDiscovery(endpoints)
	defer ser.Close()
	ser.WatchService("/web/")
	ser.WatchService("/gRPC/")
	log.Println(ser.GetService("/web/backend"))
	// for {
	// 	select {
	// 	case <-time.Tick(10 * time.Second):
	// 		log.Println(ser.GetServices())
	// 		log.Println(ser.GetService("/web/backend"))
	// 	}
	// }
}
