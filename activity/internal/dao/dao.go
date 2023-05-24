package dao

//
//import (
//	"context"
//	"time"
//
//	v8 "github.com/go-redis/redis/v8"
//	"github.com/google/wire"
//	"xy3-proto/pkg/conf/paladin"
//	xtime "xy3-proto/pkg/time"
//)
//
//var ProviderSet = wire.NewSet(New, NewRedis)
//
//var _dao *dao
//
//// DefaultDao .
//func DefaultDao() Dao {
//	return _dao
//}
//
//// Dao dao interface
//type Dao interface {
//	Close()
//	Ping(ctx context.Context) (err error)
//	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
//	// Article(c context.Context, id int64) (*model.Article, error)
//
//	////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	// 战令
//
//	// 获取指定战令实际购买次数
//	GetWarMakesNum(ctx context.Context, acType string) (num int32, err error)
//	// 更新实际购买次数
//	UpdateWarMakesNum(ctx context.Context, acType string, count int64) (ok bool, err error)
//
//	// 获取指定战令虚拟购买次数
//	GetWarMakesShamNum(ctx context.Context, acType string) (num int32, err error)
//	// 更新虚拟购买次数
//	UpdateWarMakesShamNum(ctx context.Context, acType string, count int64) (ok bool, err error)
//
//	// GetWarMakesAddSham 目标虚拟次数 .
//	GetWarMakesAddSham(ctx context.Context, acType string, day int32) (num int32, err error)
//	// UpdateWarMakesAddSham 更新目标虚拟次数 .
//	UpdateWarMakesAddSham(ctx context.Context, acType string, day int32, count int64) (err error)
//
//	// GetWarMakesIsAdd 获取指定天数指定活动是否增加了虚拟次数 .
//	GetWarMakesIsAdd(ctx context.Context, acType string, day int32) (num int32, err error)
//	// UpdateWarMakesIsAdd 记录指定天数增加的虚拟购买次数 .
//	UpdateWarMakesIsAdd(ctx context.Context, acType string, day int32, count int64) (err error)
//}
//
//// dao dao.
//type dao struct {
//	redis      *v8.Client
//	demoExpire int32
//}
//
//// New new a dao and return.
//func New(r *v8.Client) (d Dao, cf func(), err error) {
//	return newDao(r)
//}
//
//func newDao(r *v8.Client) (d *dao, cf func(), err error) {
//	var cfg struct {
//		DemoExpire xtime.Duration
//	}
//	if err = paladin.Get("application.txt").UnmarshalTOML(&cfg); err != nil {
//		return
//	}
//	d = &dao{
//		redis:      r,
//		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
//	}
//	cf = d.Close
//	_dao = d
//	return
//}
//
//// Close close the resource.
//func (d *dao) Close() {
//}
//
//// Ping ping the resource.
//func (d *dao) Ping(ctx context.Context) (err error) {
//	return nil
//}
