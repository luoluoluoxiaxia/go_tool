package easytime

import (
	"time"
)

var YmdHislayout = "2006-01-02 15:04:05"
var DateY = "2006"
var DateM = "01"
var DateD = "02"
var DateH = "15"
var DateI = "04"
var DateS = "05"
var DefaultDateTimeZone = "Asia/Shanghai"

/*
	date 时间
	layouts 时间格式
	dateTimeZone 时区
	return 根据格式返回对应时区的时间戳
*/
func DateToUnix(date string, layouts string, dateTimeZone string) int64 {

	if layouts == "" {
		layouts = YmdHislayout
	}

	if dateTimeZone == "" {
		dateTimeZone = DefaultDateTimeZone
	}

	loc, _ := time.LoadLocation(dateTimeZone)
	tmp, _ := time.ParseInLocation(layouts, date, loc)
	unix := tmp.Unix()

	return unix

}

/*
	unix 时间戳
	layouts 时间格式
	dateTimeZone 时区
	return 根据时间戳返回对应格式的时间
*/
func UnixToDate(unix int64, layouts string, dateTimeZone string) string {

	if layouts == "" {

		layouts = YmdHislayout

	}

	if dateTimeZone == "" {
		dateTimeZone = DefaultDateTimeZone
	}

	loc, _ := time.LoadLocation(dateTimeZone)

	date := time.Unix(unix, 0).In(loc).Format(layouts)

	return date

}

/*
	layouts 时间格式
	dateTimeZone 时区
	return 返回对应格式的当前时区的时间
*/

func ToDate(layouts string, dateTimeZone string) (string, error) {

	if layouts == "" {
		layouts = YmdHislayout
	}

	if dateTimeZone == "" {
		dateTimeZone = DefaultDateTimeZone
	}

	local, err := time.LoadLocation(dateTimeZone)

	if err != nil {
		return "", err

	}

	return time.Now().In(local).Format(layouts), nil
}

/*
	dateTimeZone 时区
	return 返回对应时区当前时间戳
*/
func ToUnix(dateTimeZone string) int64 {

	if dateTimeZone == "" {
		dateTimeZone = DefaultDateTimeZone
	}

	loc, _ := time.LoadLocation(dateTimeZone)

	return time.Now().In(loc).Unix()

}
