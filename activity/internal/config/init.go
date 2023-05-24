// 初始化配置 - 每日，每周，活跃度

package config

func init() {
	mm["activity_war_makes_sham.json"] = shamCfg
	mm["activity_open.json"] = openCfg

	// xcfg.Panic(xcfg.Watch("task_week_activity.json", taskWeekActiveCfg))
	// xcfg.Panic(xcfg.Watch("task_week.json", taskWeekCfg))
	// xcfg.Panic(xcfg.Watch("task_daily_control.json", taskDailyControlCfg))
	// xcfg.Panic(xcfg.Watch("task_week_control.json", taskWeekControlCfg))
}
