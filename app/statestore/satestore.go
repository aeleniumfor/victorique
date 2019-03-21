package statestore

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func SetKey() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5)
	resp, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	if err != nil {
		// handle error!
	}
	fmt.Println(resp)

	defer cli.Close()
}
