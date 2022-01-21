package cache

import (
	"context"
	"log"
	"server/cache/rpc"

	"google.golang.org/grpc"
)

type CacheClient struct {
	rpc.CacheClient
}

func NewCacheServer() *CacheClient {
	return &CacheClient{}
}
func (cc *CacheClient) Get(key string) (value string, err error) {
	request := &rpc.GetRequest{
		Key: key,
	}

	res, err := cc.CacheClient.Get(context.Background(), request)
	if err != nil {
		return
	}
	value = res.Value

	return
}

func (cc *CacheClient) Put(key string, value string) (err error) {
	request := &rpc.PutRequest{
		Key:   key,
		Value: value,
	}

	_, err = cc.CacheClient.Put(context.Background(), request)

	return
}

func NewCacheClient(addr string) *CacheClient {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(addr, opts)
	if err != nil {
		log.Fatal(err)
	}
	//defer cc.Close()

	client := rpc.NewCacheClient(cc)

	return &CacheClient{
		client,
	}
}
