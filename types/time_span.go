package types

import (
	"errors"
)

var millisecondsPerSecond = NewNumber(1000)
var hoursPerDay = NewNumber(24)
var minutesPerDay = hoursPerDay.Multiply(*NewNumber(60))
var secondsPerDay = minutesPerDay.Multiply(*NewNumber(60))
var millisecondsPerDay = secondsPerDay.Multiply(*millisecondsPerSecond)
var minutesPerHour = NewNumber(60)
var secondsPerHour = minutesPerHour.Multiply(*NewNumber(60))
var millisecondsPerHour = secondsPerHour.Multiply(*millisecondsPerSecond)
var secondsPerMinute = secondsPerHour.Divide(*NewNumber(60))
var millisecondsPerMinute = secondsPerMinute.Multiply(*millisecondsPerSecond)

type TimeSpan struct {
	day                   int64
	hour                  int64
	minute                int64
	second                int64
	millisecond           int64
	isValid               bool
	MillisecondsPerSecond *Number
	HoursPerDay           *Number
	MinutesPerDay         *Number
	SecondsPerDay         *Number
	MillisecondsPerDay    *Number
	MinutesPerHour        *Number
	SecondsPerHour        *Number
	MillisecondsPerHour   *Number
	SecondsPerMinute      *Number
	MillisecondsPerMinute *Number
}

func EmptyTimeSpan() *TimeSpan {
	return &TimeSpan{
		day:                   0,
		hour:                  0,
		minute:                0,
		second:                0,
		millisecond:           0,
		isValid:               false,
		MillisecondsPerSecond: millisecondsPerSecond,
		HoursPerDay:           hoursPerDay,
		MinutesPerDay:         minutesPerDay,
		SecondsPerDay:         secondsPerDay,
		MillisecondsPerDay:    millisecondsPerDay,
		MinutesPerHour:        minutesPerHour,
		SecondsPerHour:        secondsPerHour,
		MillisecondsPerHour:   millisecondsPerHour,
		SecondsPerMinute:      secondsPerMinute,
		MillisecondsPerMinute: millisecondsPerMinute,
	}
}

func NewTimeSpan(days, hours, minutes, seconds, milliseconds Number) *TimeSpan {
	return &TimeSpan{
		day:                   days.AsInt64(),
		hour:                  hours.AsInt64(),
		minute:                minutes.AsInt64(),
		second:                seconds.AsInt64(),
		millisecond:           milliseconds.AsInt64(),
		isValid:               true,
		MillisecondsPerSecond: millisecondsPerSecond,
		HoursPerDay:           hoursPerDay,
		MinutesPerDay:         minutesPerDay,
		SecondsPerDay:         secondsPerDay,
		MillisecondsPerDay:    millisecondsPerDay,
		MinutesPerHour:        minutesPerHour,
		SecondsPerHour:        secondsPerHour,
		MillisecondsPerHour:   millisecondsPerHour,
		SecondsPerMinute:      secondsPerMinute,
		MillisecondsPerMinute: millisecondsPerMinute,
	}
}

func NewTimeSpanFromMilliseconds(milliseconds Number) *TimeSpan {
	d, h, m, s, mi := readMilliseconds(milliseconds)
	return &TimeSpan{
		day:                   d,
		hour:                  h,
		minute:                m,
		second:                s,
		millisecond:           mi,
		isValid:               true,
		MillisecondsPerSecond: millisecondsPerSecond,
		HoursPerDay:           hoursPerDay,
		MinutesPerDay:         minutesPerDay,
		SecondsPerDay:         secondsPerDay,
		MillisecondsPerDay:    millisecondsPerDay,
		MinutesPerHour:        minutesPerHour,
		SecondsPerHour:        secondsPerHour,
		MillisecondsPerHour:   millisecondsPerHour,
		SecondsPerMinute:      secondsPerMinute,
		MillisecondsPerMinute: millisecondsPerMinute,
	}
}

// create a new Time Span from the ISO String "Day.Hour:Minute:Second Millisecond"
// values that not in the String was defined with 0
func NewTimeSpanFromISOString(isoString String) (*TimeSpan, error) {
	t, err := isoStringReader(isoString)
	if err != nil {
		return EmptyTimeSpan(), errors.New("Error on TimeSpan NewTimeSpanFromISOString: " + err.Error())
	}
	return t, nil
}

func (ts *TimeSpan) IsValid() bool {
	return ts.isValid
}

func (ts *TimeSpan) AsString() *String {
	return isoStringWriter(*ts)
}

func (ts *TimeSpan) Day() *Number {
	return NewNumber(ts.day)
}
func (ts *TimeSpan) Hour() *Number {
	return NewNumber(ts.hour)
}
func (ts *TimeSpan) Minute() *Number {
	return NewNumber(ts.minute)
}
func (ts *TimeSpan) Second() *Number {
	return NewNumber(ts.second)
}
func (ts *TimeSpan) Millisecond() *Number {
	return NewNumber(ts.millisecond)
}

/*
func (ts *TimeSpan) TotalDays() *Number         {}
func (ts *TimeSpan) TotalHours() *Number        {}
func (ts *TimeSpan) TotalMinutes() *Number      {}
func (ts *TimeSpan) TotalSeconds() *Number      {}
func (ts *TimeSpan) TotalMilliseconds() *Number {}

func (ts *TimeSpan) IsBefore(duration TimeSpan) bool      {}
func (ts *TimeSpan) IsAfter(duration TimeSpan) bool       {}
func (ts *TimeSpan) Negate() *TimeSpan                    {}
func (ts *TimeSpan) Add(duration TimeSpan) *TimeSpan      {}
func (ts *TimeSpan) Subtract(duration TimeSpan) *TimeSpan {}
func (ts *TimeSpan) Equals(duration TimeSpan) bool        {}
*/

func isoStringReader(value String) (*TimeSpan, error) {
	daySeperator := NewString(".")
	timeSeperator := NewString(":")
	msSeperator := NewString(" ")

	isInLengthBorders := value.Length().IsAbove(*NewNumber(1)) && value.Length().IsBelow(*NewNumber(15))
	daySeperatorExists := value.ContainsCount(*daySeperator).Equals(*NewNumber(1))
	timeSeperatorExists := value.ContainsCount(*timeSeperator).Equals(*NewNumber(2))
	msSeperatorExists := value.ContainsCount(*msSeperator).Equals(*NewNumber(1))

	if !isInLengthBorders {
		println("Error on TimeSpan isoStringReader: string is to short or to long")
		return EmptyTimeSpan(), errors.New("string is to short or to long")
	}

	var daySplit []*String
	if daySeperatorExists {
		daySplit = value.Split(*daySeperator)
	}
	var timeSplit []*String
	if timeSeperatorExists {
		if daySplit != nil && len(daySplit) > 1 {
			timeSplit = daySplit[1].Split(*timeSeperator)
		} else {
			timeSplit = value.Split(*timeSeperator)
		}
	}
	var msSplit []*String
	if msSeperatorExists {
		if timeSplit != nil && len(timeSplit) > 2 {
			msSplit = timeSplit[2].Split(*msSeperator)
		} else {
			msSplit = value.Split(*msSeperator)
		}
	}
	println(msSplit)

	day := 0
	hour := 0
	minute := 0
	second := 0
	millisecond := 0
	if daySplit != nil && len(daySplit) > 0 {
		day = daySplit[0].AsNumber().AsInt()
	}
	if timeSplit != nil && len(timeSplit) > 2 {
		hour = timeSplit[0].AsNumber().AsInt()
		minute = timeSplit[1].AsNumber().AsInt()
		if msSplit == nil || len(msSplit) <= 1 {
			second = timeSplit[2].AsNumber().AsInt()
		}
	}
	if msSplit != nil && len(msSplit) > 1 {
		second = msSplit[0].AsNumber().AsInt()
		millisecond = msSplit[1].AsNumber().AsInt()
	}

	return NewTimeSpan(
		*NewNumber(day),
		*NewNumber(hour),
		*NewNumber(minute),
		*NewNumber(second),
		*NewNumber(millisecond),
	), nil
}

func isoStringWriter(value TimeSpan) *String {
	return EmptyString()
}

func readMilliseconds(milliseconds Number) (day, hour, minute, second, millisecond int64) {
	dayNumber := milliseconds.Divide(*millisecondsPerDay).Floor(*NewNumber(0))
	hourNumber := milliseconds.Divide(*millisecondsPerHour).Floor(*NewNumber(0)).
		Subtract(*dayNumber.Multiply(*millisecondsPerDay))
	minuteNumber := milliseconds.Divide(*millisecondsPerMinute).Floor(*NewNumber(0)).
		Subtract(*dayNumber.Multiply(*minutesPerDay)).
		Subtract(*hourNumber.Multiply(*minutesPerHour))
	secondNumber := milliseconds.Divide(*millisecondsPerSecond).Floor(*NewNumber(0)).
		Subtract(*dayNumber.Multiply(*secondsPerDay)).
		Subtract(*hourNumber.Multiply(*secondsPerHour)).
		Subtract(*minuteNumber.Multiply(*secondsPerMinute))
	millisecondNumber := milliseconds.
		Subtract(*dayNumber.Multiply(*millisecondsPerDay)).
		Subtract(*hourNumber.Multiply(*millisecondsPerHour)).
		Subtract(*minuteNumber.Multiply(*millisecondsPerMinute)).
		Subtract(*secondNumber.Multiply(*millisecondsPerSecond))
	return dayNumber.AsInt64(),
		hourNumber.AsInt64(),
		minuteNumber.AsInt64(),
		secondNumber.AsInt64(),
		millisecondNumber.AsInt64()
}
