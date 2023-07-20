package etcd

import (
	"log"
	"testing"
)

func TestRegistry(t *testing.T) {
	var endpoints = []string{"localhost:2379"}
	_, err := NewServiceRegister(endpoints, "/web/backend", "localhost:8000", 5)
	if err != nil {
		log.Fatalln(err)
	}
	//监听续租相应chan
	// go svc.ListenLeaseRespChan()
	// select {
	// // case <-time.After(20 * time.Second):
	// // 	ser.Close()
	// }
}
