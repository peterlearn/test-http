// 活动启动开关配置

package config

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

var (
	openCfg = &openConfig{}
)

// OpenConfig 活动开关配置集合 .
type openConfig struct {
	items map[int32]*Open
	hash  string // 根据配置生成的hash值
}

// Open 活动开关的详细配置 .
type Open struct {
	Order       int32  `json:"order"`       // 序号
	ActivityID  int32  `json:"activityID"`  // 活动ID
	Deadline    int32  `json:"deadline"`    // 活动期限
	PutawayTime string `json:"putawayTime"` // 上架时间
	SoldOutTime string `json:"soldOutTime"` // 下架时间
	StartTime   string `json:"startTime"`   // 公示开始时间
	EndTime     string `json:"endTime"`     // 公示结束时间
	IsOpen      int32  `json:"isOpen"`      // 是否开启
	Name        string `json:"name"`        // 活动名称
}

// Set .
func (c *openConfig) Set(text []byte) error {
	var ec = openConfig{
		items: make(map[int32]*Open),
	}
	var data []*Open
	if err := json.Unmarshal(text, &data); err != nil {
		return err
	}
	sha := sha256.New()
	sha.Write(text)
	bytes := sha.Sum(nil)
	ec.hash = hex.EncodeToString(bytes)

	for _, v := range data {
		ec.items[v.ActivityID] = v
	}

	*c = ec
	return nil
}

// GetOpenConfig 根据活动ID获取指定的配置 .
func GetOpenConfig(id int32) *Open {
	if v, ok := openCfg.items[id]; ok {
		return v
	}
	return nil
}

//  GetAllOpenConfig 获取所有配置项目 .
func GetAllOpenConfig() map[int32]*Open {
	return openCfg.items
}

// GetOpenConfigHash 获取当前配置生成的hash .
func GetOpenConfigHash() string {
	return openCfg.hash
}

// OpenConfigHashIsEqual 校验hash 是否一致 .
func OpenConfigHashIsEqual(h string) bool {
	return openCfg.hash == h
}
