// 活动相关的redis操作

package dao

//import (
//	"context"
//	"x-server/activity/internal/model"
//	"x-server/core/pkg/util"
//)
//
//// GetWarMakesNum 获取指定战令实际购买次数 .
//func (d *dao) GetWarMakesNum(ctx context.Context, acType string) (num int32, err error) {
//	key := model.KeyActivityWarMakes()
//	val, err := d.redis.HGet(context.TODO(), key, acType).Result()
//	if err != nil {
//		return
//	}
//	num = util.StrToInt32(val)
//	return
//}
//
//// UpdateWarMakesNum 更新实际购买次数 .
//func (d *dao) UpdateWarMakesNum(ctx context.Context, acType string, count int64) (ok bool, err error) {
//	key := model.KeyActivityWarMakes()
//	val, err := d.redis.HIncrBy(context.TODO(), key, acType, count).Result()
//	if err != nil {
//		return
//	}
//	if val < 0 {
//		return false, nil
//	}
//	return true, nil
//}
//
//// GetWarMakesShamNum 获取指定战令虚拟购买次数 .
//func (d *dao) GetWarMakesShamNum(ctx context.Context, acType string) (num int32, err error) {
//	key := model.KeyActivityWarMakesSham()
//	val, err := d.redis.HGet(context.TODO(), key, acType).Result()
//	if err != nil {
//		return
//	}
//	num = util.StrToInt32(val)
//	return
//}
//
//// UpdateWarMakesShamNum 更新虚拟购买次数 .
//func (d *dao) UpdateWarMakesShamNum(ctx context.Context, acType string, count int64) (ok bool, err error) {
//	key := model.KeyActivityWarMakesSham()
//	val, err := d.redis.HIncrBy(context.TODO(), key, acType, count).Result()
//	if err != nil {
//		return
//	}
//	if val < 0 {
//		return false, nil
//	}
//	return true, nil
//}
//
//// GetWarMakesAddSham 获取要增加的虚拟购买次数 .
//func (d *dao) GetWarMakesAddSham(ctx context.Context, acType string, day int32) (num int32, err error) {
//	key := model.KeyActivityWarMakesAddSham(day)
//	val, err := d.redis.HGet(context.TODO(), key, acType).Result()
//	if err != nil {
//		return
//	}
//	num = util.StrToInt32(val)
//	return
//}
//
//// UpdateWarMakesAddSham 更新剩余要增加的虚拟次数 .
//func (d *dao) UpdateWarMakesAddSham(ctx context.Context, acType string, day int32, count int64) (err error) {
//	key := model.KeyActivityWarMakesAddSham(day)
//	_, err = d.redis.HIncrBy(context.TODO(), key, acType, count).Result()
//	return
//}
//
//// GetWarMakesIsAdd 获取指定天数指定活动是否增加了虚拟次数 .
//func (d *dao) GetWarMakesIsAdd(ctx context.Context, acType string, day int32) (num int32, err error) {
//	key := model.KeyActivityWarMakesIsAdd(day)
//	val, err := d.redis.HGet(context.TODO(), key, acType).Result()
//	if err != nil {
//		return
//	}
//	num = util.StrToInt32(val)
//	return
//}
//
//// UpdateWarMakesIsAdd 记录指定天数增加的虚拟购买次数 .
//func (d *dao) UpdateWarMakesIsAdd(ctx context.Context, acType string, day int32, count int64) (err error) {
//	key := model.KeyActivityWarMakesIsAdd(day)
//	_, err = d.redis.HIncrBy(context.TODO(), key, acType, count).Result()
//	return
//}
