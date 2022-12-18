package helpers

import (
	"context"
	"log"

	globals "backend/globals"

	"go.etcd.io/etcd/clientv3"
)

func GetKeyFromRaft(key string) *clientv3.GetResponse {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   globals.Endpoints,
		DialTimeout: globals.Timeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx2, cancel := context.WithTimeout(context.Background(), globals.Timeout)
	resp, err := cli.Get(ctx2, key)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func PutValueForKeys(key string, value string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   globals.Endpoints,
		DialTimeout: globals.Timeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	_, err = cli.Put(context.TODO(), key, value)
	if err != nil {
		log.Fatal(err)
	}
}
