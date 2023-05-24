package model

//import (
//	"flag"
//	"os"
//	"strconv"
//
//	pbactivity "xy3-proto/activity"
//
//	"xy3-proto/pkg/conf/env"
//)
//
//var (
//	ServiceName    string
//	ServiceID      string
//	ServiceVersion int64
//	ZoneID         int64 // 战区id
//	ServerID       int64 // 服id
//)
//
//func init() {
//	addFlag(flag.CommandLine)
//
//	ServiceName = pbactivity.ServiceName
//
//	k8sFlag := os.Getenv("KUBERNETES_SERVICE_HOST")
//	if len(k8sFlag) > 0 {
//		// StatefulSet 用os.Hostname();  Deployment 用pb.ServiceName
//		ServiceID = pbactivity.ServiceName
//		//model.ServiceID, _ = os.Hostname()
//	}
//	if os.Getenv("DISCOVERY") != "" {
//		ServiceID = env.AppID
//	}
//}
//
//func addFlag(fs *flag.FlagSet) {
//	v := os.Getenv("DISCOVERY_GRPC_ADDR")
//
//	v = os.Getenv("VERSION")
//	version, _ := strconv.ParseInt(v, 10, 64)
//	fs.Int64Var(&ServiceVersion, "version", version, "service version , default: 0")
//	v = os.Getenv("ZONEID")
//	zoneid, _ := strconv.ParseInt(v, 10, 64)
//	fs.Int64Var(&ZoneID, "zoneid", zoneid, "zoneid , default: 0")
//	v = os.Getenv("SERVERID")
//	serverid, _ := strconv.ParseInt(v, 10, 64)
//	fs.Int64Var(&ServerID, "serverid", serverid, "serverid , default: 0")
//}
