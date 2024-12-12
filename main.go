package moment

import "time"

// IsLeapYear 是否为闰年
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

// MonthDiffer 获取两个时间之间的月份差。精确到月。例如在同年同月，返回0。结果可为负。
func MonthDiffer(after time.Time, before time.Time) int {
	give := true
	if after.Before(before) {
		after, before = before, after
		give = false
	}
	sub := (after.Year()-before.Year())*12 + int(after.Month()) - int(before.Month())
	if !give {
		return 0 - sub
	}
	return sub
}

// AddDateByMonth 对time.AddDate的本地化封装。以月为精度相加。例如1月31日加1个月，返回2月28日，如为闰年，返回2月29日
func AddDateByMonth(t time.Time, year int, month int) time.Time {
	y, m, d := t.Date()
	hour, mi, sec := t.Clock()
	loc := t.Location()
	var ySub, mSub int
	month = int(m) + month
	if month == 0 {
		ySub = 0
		mSub = 0
	} else if month%12 == 0 && month < 0 {
		mSub = 12
		ySub = month/12 + 1
	} else if month%12 == 0 && month > 0 {
		mSub = 12
		ySub = month/12 - 1
	} else if month < 0 {
		ySub = month/12 - 1
		mSub = month%12 + 12
	} else {
		ySub = month / 12
		mSub = month % 12
	}
	year = y + ySub + year
	if mSub == 2 && IsLeapYear(year) && d > 29 {
		d = 29
	} else if mSub == 2 && !IsLeapYear(year) && d > 28 {
		d = 28
	}
	return GetCorrectDate(year, mSub, d, hour, mi, sec, 0, loc)
}

// GetCorrectDate 对time.Date的本地化封装。如果日期过大，直接得到当月的最后一天。例如闰年2月31日，得到2月29日
func GetCorrectDate(year int, month int, day int, hour int, min int, sec int, nsec int, loc *time.Location) time.Time {
	if day < 1 {
		return time.Date(year, time.Month(month), 1, hour, min, sec, nsec, loc)
	}
	switch month {
	case 2:
		if IsLeapYear(year) && day > 29 {
			day = 29
		} else if !IsLeapYear(year) && day > 28 {
			day = 28
		}
	case 4, 6, 9, 11:
		if day > 30 {
			day = 30
		}
	default:
		if day > 31 {
			day = 31
		}
	}
	return time.Date(year, time.Month(month), day, hour, min, sec, nsec, loc)
}
