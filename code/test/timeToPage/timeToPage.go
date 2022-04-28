package timeToPage

import "time"

// 传递 开始时间/结束时间 分页 和 每页数量
// 需求: 通过分页, 从新计算出 开始时间到截止时间的时间范围, 减少查询的数量
const PageSize = 10

type StatisticsPage struct {
	StartTime time.Time
	EndTime   time.Time
	Total     int64
	Dates     []time.Time
}

func ResetTimeAccordingToPage(startTimeReq int64, endTimeReq int64, page int64, pageSize int64) (res StatisticsPage) {
	startUnix := time.Unix(startTimeReq, 0)
	startDay := time.Date(startUnix.Year(), startUnix.Month(), startUnix.Day(), 0, 0, 0, 0, time.Local)

	endUnix := time.Unix(endTimeReq, 0)
	endDay := time.Date(endUnix.Year(), endUnix.Month(), endUnix.Day(), 0, 0, 0, 0, time.Local)
	dates := []time.Time{}
	//如果每页分页，返回所有
	if page < 1 {

		res.EndTime = time.Date(endDay.Year(), endDay.Month(), endDay.Day(), 59, 59, 59, 0, time.Local)
		res.StartTime = startDay
		for {
			dates = append(dates, endDay)
			endDay = endDay.Add(-24 * time.Hour)
			if endDay.Unix() < startDay.Unix() {
				break
			}
		}
		res.Total = int64(len(dates))
		res.Dates = dates
		return
	}
	//根据日期返回对应的开始结束日期

	diffDay := int64(endDay.Sub(startDay).Hours() / 24)
	res.Total = diffDay + 1

	//算出开始日期
	if pageSize <= 0 {
		pageSize = int64(PageSize) // 默认的pageSize
	}
	offset := (page - 1) * pageSize
	offsetDay := endDay.Add(-24 * time.Duration(offset) * time.Hour)
	//算出长度
	var i int64
	for i = 0; i < pageSize; i++ {
		dates = append(dates, offsetDay)
		offsetDay = offsetDay.Add(-24 * time.Hour)
		if offsetDay.Unix() < startDay.Unix() {
			break
		}
	}
	res.StartTime = startDay
	res.EndTime = time.Date(endDay.Year(), endDay.Month(), endDay.Day(), 59, 59, 59, 0, time.Local)
	res.Dates = dates

	return
}
