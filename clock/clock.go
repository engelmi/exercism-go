package clock

import (
	"fmt"
	"strconv"
)

type Clock struct {
	hour   int
	minute int
}

func trim(c *Clock) {
	if c.minute > 59 {
		mm := c.minute / 60
		mr := c.minute % 60
		c.hour += mm
		c.minute = mr
	}
	if c.minute < 0 {
		mm := c.minute / 60
		c.hour += mm

		mr := c.minute - mm*60
		c.minute = 60 + mr
		if mr < 0 {
			c.hour -= 1
		}
	}
	c.hour = c.hour % 24
	c.minute = c.minute % 60

	if c.hour > 23 {
		c.hour = c.hour % 24
	}
	if c.hour < 0 {
		c.hour = 24 - ((c.hour * -1) % 24)
	}
}

func New(h int, m int) Clock {
	c := Clock{
		hour:   h,
		minute: m,
	}
	trim(&c)

	return c
}

func (c Clock) String() string {
	hStr := strconv.Itoa(c.hour)
	if c.hour < 10 {
		hStr = "0" + strconv.Itoa(c.hour)
	}
	mStr := strconv.Itoa(c.minute)
	if c.minute < 10 {
		mStr = "0" + strconv.Itoa(c.minute)
	}
	return fmt.Sprintf("%s:%s", hStr, mStr)
}

func (c Clock) Add(m int) Clock {
	c.minute += m
	trim(&c)

	return c
}

func (c Clock) Subtract(m int) Clock {
	c.minute -= m
	trim(&c)

	return c
}
