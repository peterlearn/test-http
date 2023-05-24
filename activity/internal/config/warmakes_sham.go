// 战令 虚拟次数增加配置

package config

import "encoding/json"

var (
	shamCfg = &shamConfig{}
)

// shamConfig 活动开关配置集合 .
type shamConfig struct {
	items map[int32]*Sham
}

// Sham 活动开关的详细配置 .
type Sham struct {
	ID         int32    `json:"id"`         // 唯一索引
	ActivityID int32    `json:"activityID"` // 活动ID
	Day        int32    `json:"day"`        // 活动开启天数
	Section    [2]int32 `json:"section"`    // 虚拟次数区间段
}

// Set .
func (c *shamConfig) Set(text []byte) error {
	var ec = shamConfig{
		items: make(map[int32]*Sham),
	}
	var data []*Sham
	if err := json.Unmarshal(text, &data); err != nil {
		return err
	}
	for _, v := range data {
		ec.items[v.ID] = v
	}

	*c = ec
	return nil
}

// GetWarMakesConfig 根据活动id和活动开启天数
// id 活动ID
// day 当天活动开启的天数
func GetWarMakesConfig(id int32, day int32) *Sham {
	for _, v := range shamCfg.items {
		if v.ActivityID != id {
			continue
		}
		if v.Day != day {
			continue
		}
		return v
	}
	return nil
}

// GetExistWarMakesConfig 获取最近存在的配置 .
func GetExistWarMakesConfig(id int32, day int32) *Sham {
	maxDay := int32(0)
	for _, v := range shamCfg.items {
		if id != v.ActivityID {
			continue
		}
		if day <= v.Day {
			continue
		}
		if maxDay < v.Day {
			maxDay = v.Day
		}
	}
	return GetWarMakesConfig(id, maxDay)
}

// 获取指定天数的下一个配置
// func GetNextWarMakesConfig(id int32, day int32) *Sham  {

// }
