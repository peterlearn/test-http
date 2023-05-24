// 常量 - 声明定义

package model

//import (
//	"fmt"
//
//	env "xy3-proto/pkg/conf/env"
//)
//
//const (
//	REDIS_ACTIVITY_WARMAKES_WELFARE_INFO    = "activity:warmakes:server:%v:info"           // 不同战令的实际全民购买次数(hashmap) field:活动 value：购买次数
//	REDIS_ACTIVITY_WARMAKES_WELFARE_SHAM    = "activity:warmakes:server:%v:shaminfo"       // 虚拟购买次数(hashmap) field:活动 value：购买次数
//	REDIS_ACTIVITY_WARMAKES_WELFARE_ADDSHAM = "activity:warmakes:server:%v:day:%v:addsham" // 记录指定天数内要增加到的购买次数(hashmap) field:活动 value：要增加的虚拟次数
//	REDIS_ACTIVITY_WARMAKES_WELFARE_ISADD   = "activity:warmakes:server:%v:day:%v:isadd"   // 记录指定天数是否增加了虚拟次数(hashmap) field:活动 value：当天增加的次数
//
//	WARMAKES_WELFARE_BUYNUM = 1
//)
//
//// KeyActivityWarMakes 活动实际购买次数 .
//func KeyActivityWarMakes() string {
//	return fmt.Sprintf(REDIS_ACTIVITY_WARMAKES_WELFARE_INFO, env.Namespace)
//}
//
//// KeyActivityWarMakesSham 活动虚拟购买次数 .
//func KeyActivityWarMakesSham() string {
//	return fmt.Sprintf(REDIS_ACTIVITY_WARMAKES_WELFARE_SHAM, env.Namespace)
//}
//
//// KeyActivityWarMakesAddSham 记录要增加虚拟购买次数 .
//func KeyActivityWarMakesAddSham(day int32) string {
//	return fmt.Sprintf(REDIS_ACTIVITY_WARMAKES_WELFARE_ADDSHAM, env.Namespace, day)
//}
//
//// KeyActivityWarMakesIsAdd 活动虚拟购买次数 .
//func KeyActivityWarMakesIsAdd(day int32) string {
//	return fmt.Sprintf(REDIS_ACTIVITY_WARMAKES_WELFARE_ISADD, env.Namespace, day)
//}
