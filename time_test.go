package moment

import (
	"fmt"
	"testing"
	"time"
)

type Person struct {
	Birthday  time.Time
	RetireDay time.Time
	Gender    string
}

func (p *Person) setRetireDay() {
	var base int
	var start time.Time
	if p.Gender == "男" {
		base = 60
		start = time.Date(1965, 1, 1, 0, 0, 0, 0, time.UTC)
	} else {
		base = 55
		start = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	sub := MonthDiffer(p.Birthday, start)
	if sub < 0 {
		p.RetireDay = p.Birthday.AddDate(base, 0, 0)
		return
	}
	p.RetireDay = p.Birthday.AddDate(base, sub/4+1, 0)
}

func TestTime(t *testing.T) {
	for i := 1963; i <= 1981; i++ {
		for j := 1; j <= 12; j++ {
			p := Person{
				Birthday:  time.Date(i, time.Month(j), 2, 0, 0, 0, 0, time.UTC),
				RetireDay: time.Time{},
				Gender:    "女",
			}
			p.setRetireDay()
			fmt.Println(p)
		}
	}
}

func TestAdd(t *testing.T) {
	b := time.Date(2002, 4, 30, 12, 0, 0, 0, time.Local)
	fmt.Println(b.AddDate(-2, -2, 0))
	fmt.Println(SimpleAdd(b, -1, -2))
}
func TestA(t *testing.T) {
	fmt.Println(-32 % 12)
}
