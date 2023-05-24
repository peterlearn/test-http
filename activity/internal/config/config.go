package config

import (
	"github.com/go-kratos/kratos/v2/log"
	"gitlab.okymo.com.tw/middle-end/kratos/pkg/conf/paladin"
)

var (
	mm = make(map[string]paladin.Setter)
)

func Init() error {
	if err := paladin.Init(); err != nil {
		log.Error("fileInit paladin err:%v", err)
		return err
	}
	paladin.Watch([]string{}, mm)
	return nil
}
