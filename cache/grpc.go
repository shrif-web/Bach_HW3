package main

import (
	"cache/rpc"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type CacheServer struct {
	cache *Cache
}

func NewCacheServer(capacity int) *CacheServer {
	cache := NewCache(capacity)
	return &CacheServer{
		cache: &cache,
	}
}
func (cs *CacheServer) Get(ctx context.Context, req *rpc.GetRequest) (res *rpc.Response, err error) {
	res = &rpc.Response{}
	res.Value, err = cs.cache.Get(req.Key)
	return
}

func (cs *CacheServer) Put(ctx context.Context, req *rpc.PutRequest) (res *rpc.Response, err error) {
	res = &rpc.Response{}
	err = cs.cache.Put(req.Key, req.Value)
	return
}

func (cs *CacheServer) Remove(ctx context.Context, req *rpc.GetRequest) (res *rpc.Response, err error) {
	res = &rpc.Response{}
	res.Value, err = cs.cache.Remove(req.Key)
	return
}

func (cs *CacheServer) Start(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on port %v ...", port)

	s := grpc.NewServer()
	rpc.RegisterCacheServer(s, cs)
	s.Serve(lis)
}
