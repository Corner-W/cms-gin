package timeutil

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TimeStr string

func (t TimeStr) String() string {
	return string(t)
}

type DateStr string

func (d DateStr) String() string {
	return string(d)
}

const (
	FormatDateTime     = "2006-01-02 15:04:05"
	FormatDateHyphen   = "2006-01-02"
	FormatDateFullStop = "2006.01.02"
	FormatDateDigit    = "20060102"
	TimeRangeJoint     = "~"
)

// LastDayStart 获取前一日开始时间
func LastDayStart() string {
	curTime := time.Now()        //获取系统当前时间
	h := fmt.Sprintf("-%dh", 24) //减去24小时（前一天）
	dh, _ := time.ParseDuration(h)
	timeStr := curTime.Add(dh).Format(FormatDateHyphen) //系统前一天日期格式转换
	return timeStr + " 00:00:00"
}

// LastDayEnd 获取前一日结束时间
func LastDayEnd() string {
	curTime := time.Now()        //获取系统当前时间
	h := fmt.Sprintf("-%dh", 24) //减去24小时（前一天）
	dh, _ := time.ParseDuration(h)
	timeStr := curTime.Add(dh).Format(FormatDateHyphen) //系统前一天日期格式转换
	return timeStr + " 23:59:59"
}

// DayStart 20220705 -> 2022-07-05 00:00:00
func DayStart(daily int) string {
	start, _ := time.Parse(FormatDateDigit, strconv.Itoa(daily))
	return start.Format(FormatDateHyphen) + " 00:00:00"
}

// DayEnd 20220705 -> 2022-07-05 23:59:59
func DayEnd(daily int) string {
	end, _ := time.Parse(FormatDateDigit, strconv.Itoa(daily))
	return end.Format(FormatDateHyphen) + " 23:59:59"
}

// ParseDayStart 2022-07-05 -> 2022-07-05 00:00:00
func ParseDayStart(date string) string {
	return date + " 00:00:00"
}

// DailyToDateString 20220705 -> 2022-07-05
func DailyToDateString(daily int) string {
	end, _ := time.Parse(FormatDateDigit, strconv.Itoa(daily))
	return end.Format(FormatDateHyphen)
}

// ParseDayEnd 2022-07-05 -> 2022-07-05 23:59:59
func ParseDayEnd(date string) string {
	return date + " 23:59:59"
}

// GetDailyTimeSpan 获取某天的开始与结果时刻
// daily == 0 返回前一天的开始与结束时刻
// daily !=0 返回指定某一天的开始与结束时刻
func GetDailyTimeSpan(daily int) (start, end string) {
	if daily == 0 {
		start = LastDayStart()
		end = LastDayEnd()
		return
	}
	start = DayStart(daily)
	end = DayEnd(daily)
	return
}

// GetDaily 获取某一天 返回格式20220704
// daily == 0 返回前一天
// daily != 0 返回指定的天
func GetDaily(daily int) int {
	if daily == 0 {
		return LastDayInt()
	}

	return daily
}

// LastDayInt 获取前一日 20220601
func LastDayInt() (ret int) {
	return LastDayIntWithParam(-1)
}

// LastDayIntWithParam 获取前X天的日期 20220601
func LastDayIntWithParam(beforeDay int) (ret int) {
	yesterday := time.Now().AddDate(0, 0, beforeDay)
	ret, _ = strconv.Atoi(yesterday.Format(FormatDateDigit))
	return
}

// ParseTime string根据格式转换为time
func ParseTime(timeStr string, format string) (time.Time, error) {
	return time.ParseInLocation(format, timeStr, time.Local)
}

// TimeStrConvDaily 将2021-11-07转换为20211107
func TimeStrConvDaily(str string) (result int) {

	showTime, _ := time.ParseInLocation(FormatDateHyphen, str, time.Local)
	result, _ = strconv.Atoi(showTime.Format(FormatDateDigit))
	return result
}

// GetTodayDaily 获取当天的日期 返回格式20220704
func GetTodayDaily() int {
	return LastDayIntWithParam(0)
}

// TimeDailyRange 取2021-11-07至2021-11-09  之间的天数 返回[20211107,20211108,20211109]
func TimeDailyRange(start string, end string) (result []int) {
	startTime, _ := time.ParseInLocation(FormatDateHyphen, start, time.Local)
	startTimeDaily, _ := strconv.Atoi(startTime.Format(FormatDateDigit))
	endTime, _ := time.ParseInLocation(FormatDateHyphen, end, time.Local)
	endTimeDaily, _ := strconv.Atoi(endTime.Format(FormatDateDigit))
	result = append(result, startTimeDaily)
	if startTimeDaily == endTimeDaily {
		return result
	}
	var day int
	for {
		day += 1
		nextTime := startTime.AddDate(0, 0, day)
		nextTimeDaily, _ := strconv.Atoi(nextTime.Format(FormatDateDigit))
		if nextTimeDaily == endTimeDaily {
			break
		} else {
			result = append(result, nextTimeDaily)
		}
	}
	result = append(result, endTimeDaily)
	return result
}

// TimeDailyRangeToStr 取2021-11-07至2021-11-09  之间的天数 返回[20211107,20211108,20211109]
func TimeDailyRangeToStr(start, end time.Time) (result []string) {
	startTimeDaily := start.Format(FormatDateHyphen)
	endTimeDaily := end.Format(FormatDateHyphen)
	result = append(result, startTimeDaily)
	if startTimeDaily == endTimeDaily {
		return result
	}
	var day int
	for {
		day += 1
		nextTime := start.AddDate(0, 0, day)
		nextTimeDaily := nextTime.Format(FormatDateHyphen)
		if nextTimeDaily == endTimeDaily {
			break
		} else {
			result = append(result, nextTimeDaily)
		}
	}
	result = append(result, endTimeDaily)
	return result
}

// UnixToTimeStr 将 时间戳 转换为 2021-11-07
func UnixToTimeStr(unixVal int64) (str string) {
	tm := time.Unix(unixVal, 0)
	return tm.Format(FormatDateHyphen)
}

// UnixToTimeSecondsStr 将 时间戳 转换为 2021-11-07 00:00:00
func UnixToTimeSecondsStr(unixVal int64) (str string) {
	tm := time.Unix(unixVal, 0)
	return tm.Format(FormatDateTime)
}

// UnixToDaily 将 时间戳 转换为 20211107
func UnixToDaily(unixVal int64) (daily int) {
	tm := time.Unix(unixVal, 0)
	showTime, _ := time.ParseInLocation(FormatDateHyphen, tm.Format(FormatDateHyphen), time.Local)
	daily, _ = strconv.Atoi(showTime.Format(FormatDateDigit))
	return daily
}

// YmdTimeToUnix 2021-11-07 转换为 时间戳
func YmdTimeToUnix(timeStr string) int64 {
	t, _ := time.ParseInLocation(FormatDateHyphen, timeStr, time.Local)
	return t.Unix()
}

// GetBeforeTime 获取n天前的秒时间戳、日期时间戳
// _day为负则代表取前几天，为正则代表取后几天，0则为今天
func GetBeforeTime(day int) (int64, string) {
	// 时区
	timeZone := time.FixedZone("CST", 8*3600) // 东八区

	// 前n天
	nowTime := time.Now().In(timeZone)
	beforeTime := nowTime.AddDate(0, 0, -day)

	// 时间转换格式
	beforeTimeS := beforeTime.Unix()                                 // 秒时间戳
	beforeDate := time.Unix(beforeTimeS, 0).Format("20060102150405") // 固定格式的日期时间戳

	return beforeTimeS, beforeDate
}

// CountWorkkdays 计算工作日
func CountWorkkdays(start time.Time, end time.Time) int {
	start = start.Local()
	end = end.Local()
	workdayCount := 0
	current := start
	for current.Before(end) || current.Equal(end) {
		if current.Weekday() != time.Saturday && current.Weekday() != time.Sunday {
			workdayCount++
		}

		current = current.AddDate(0, 0, 1)
	}

	return workdayCount
}

// GetDay 获取当天时间戳，不含时间
func GetDay(day int) time.Time {
	// 获取当前日期
	currentDate := time.Now().AddDate(0, 0, day)
	// 设置时间为 00:00:01
	return time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 1, 0, time.Local)
}

func PbtimeToDatestr(t *timestamppb.Timestamp) DateStr {
	if t == nil {
		return ""
	}
	return DateStr(t.AsTime().Local().Format(FormatDateHyphen))
}
func DatestrToPbtime(s DateStr) *timestamppb.Timestamp {
	if s == "" {
		return nil
	}
	t, _ := time.ParseInLocation(FormatDateHyphen, s.String(), time.Local)
	return timestamppb.New(t)
}
func PbtimeToTimestr(t *timestamppb.Timestamp) TimeStr {
	if t == nil {
		return ""
	}
	return TimeStr(t.AsTime().Local().Format(FormatDateTime))
}
func TimestrToPbtime(s TimeStr) *timestamppb.Timestamp {
	if s == "" {
		return nil
	}
	t, _ := time.ParseInLocation(FormatDateTime, s.String(), time.Local)
	return timestamppb.New(t)
}

func ConverStringOption() copier.Option {
	opt := copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    false,
		Converters: []copier.TypeConverter{
			{
				SrcType: &timestamppb.Timestamp{},
				DstType: DateStr(""),
				Fn: func(src any) (any, error) {
					t, ok := src.(*timestamppb.Timestamp)
					if !ok {
						return nil, errors.New("src type not matching")
					}
					return PbtimeToDatestr(t), nil
				},
			},
			{
				SrcType: &timestamppb.Timestamp{},
				DstType: TimeStr(""),
				Fn: func(src any) (any, error) {
					t, ok := src.(*timestamppb.Timestamp)
					if !ok {
						return nil, errors.New("src type not matching")
					}
					return PbtimeToTimestr(t), nil
				},
			},
		},
	}
	return opt
}

func ConverPbtimeOption() copier.Option {
	opt := copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    false,
		Converters: []copier.TypeConverter{
			{
				SrcType: DateStr(""),
				DstType: &timestamppb.Timestamp{},
				Fn: func(src any) (any, error) {
					s, ok := src.(DateStr)
					if !ok {
						return nil, errors.New("src type not matching")
					}

					return DatestrToPbtime(s), nil
				},
			},
			{
				SrcType: TimeStr(""),
				DstType: &timestamppb.Timestamp{},
				Fn: func(src any) (any, error) {
					s, ok := src.(TimeStr)
					if !ok {
						return nil, errors.New("src type not matching")
					}

					return TimestrToPbtime(s), nil
				},
			},
		},
	}
	return opt
}

// InSameDay  判断两时间是否是同一天
func InSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

// InSameWeek 判断两时间是否是同周
func InSameWeek(t1, t2 time.Time) bool {
	y1, w1 := t1.ISOWeek()
	y2, w2 := t2.ISOWeek()
	return y1 == y2 && w1 == w2
}

// InSameMonth 判断两时间是否是同月
func InSameMonth(t1, t2 time.Time) bool {
	y1, m1, _ := t1.Date()
	y2, m2, _ := t2.Date()
	return y1 == y2 && m1 == m2
}

// GetWeekStartAndEnd 获取所在周的首尾
func GetWeekStartAndEnd(t time.Time) (time.Time, time.Time) {
	weekday := t.Local().Weekday()

	var weekStart time.Time
	if weekday == time.Sunday {
		weekStart = t.AddDate(0, 0, -6)
	} else {
		daysUntilMonday := int(time.Monday - weekday)
		weekStart = t.AddDate(0, 0, daysUntilMonday)
	}

	weekEnd := weekStart.AddDate(0, 0, 6)
	return weekStart, weekEnd
}

// GetMonthStartAndEnd 获取所在月的首尾
func GetMonthStartAndEnd(t time.Time) (time.Time, time.Time) {
	monthStart := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())

	nextMonth := monthStart.AddDate(0, 1, 0)
	monthEnd := nextMonth.Add(-time.Nanosecond)

	return monthStart, monthEnd
}
