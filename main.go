package moment

import "time"

func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

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

func SimpleAdd(t time.Time, year int, month int) time.Time {
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
	return time.Date(year, time.Month(mSub), d, hour, mi, sec, 0, loc)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
