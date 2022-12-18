package helpers

import (
	"context"
	"encoding/json"
	"log"

	globals "backend/globals"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func GetValueForKey(key string) *clientv3.GetResponse {
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

func GetMap(key string) map[string][]string {
	return_map := make(map[string][]string)

	following_resp := GetValueForKey(key)
	for _, ev := range following_resp.Kvs {
		json.Unmarshal(ev.Value, &return_map)
	}

	return return_map
}

func PutMap(map_to_put map[string][]string) {
	updated_map, err := json.Marshal(map_to_put)
	if err != nil {
		log.Println(err)
	}
	PutValueForKeys("following", string(updated_map))
}
