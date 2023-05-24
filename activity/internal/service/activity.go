// 活动

package service

import (
	"context"
	pbactivity "test/activity"
	"time"
)

// ActivityVersion
// 活动版本号
func (s *Service) ActivityVersion(ctx context.Context, req *pbactivity.Empty) (resp *pbactivity.ActivityVersionResp, err error) {
	resp = &pbactivity.ActivityVersionResp{
		Version: "v1",
		Time:    time.Now().Unix(),
	}
	return
}
