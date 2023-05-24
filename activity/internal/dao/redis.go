package dao

//import (
//	"context"
//	v8 "github.com/go-redis/redis/v8"
//	"xy3-proto/pkg/conf/paladin"
//	"xy3-proto/pkg/log"
//)
//
//// NewRedis .
//func NewRedis() (r *v8.Client, cancel func(), err error) {
//	var cfg struct {
//		Client *v8.Options
//	}
//	err = paladin.Get("redis.txt").UnmarshalTOML(&cfg)
//	if err != nil {
//		log.Error("NewRedis err:%v", err)
//		return
//	}
//
//	r = v8.NewClient(cfg.Client)
//	_, err = r.Ping(context.Background()).Result()
//	if err != nil {
//		log.Error("NewRedis Ping err:%v", err)
//		return
//	}
//	cancel = func() {
//		if err := r.Close(); err != nil {
//			log.Error("NewRedis Close err:%v", err)
//			return
//		}
//	}
//
//	log.Debug("NewRedis Ping success !!!")
//	return
//}
//
//// NewRedisCluster .
//func NewRedisCluster() (r *v8.ClusterClient, err error) {
//	var cfg struct {
//		ClusterClient *v8.ClusterOptions
//	}
//	err = paladin.Get("redis.txt").UnmarshalTOML(&cfg)
//	if err != nil {
//		log.Error("NewRedisCluster err:%v", err)
//		return
//	}
//	log.Debug("NewRedisCluster redis cfg:%+v", cfg.ClusterClient)
//	r = v8.NewClusterClient(cfg.ClusterClient)
//	res, err := r.Ping(context.Background()).Result()
//	if err != nil {
//		log.Error("NewRedisCluster Ping err:%v", err)
//		return
//	}
//	log.Info("NewRedisCluster Ping Result %v", res)
//	return
//}
