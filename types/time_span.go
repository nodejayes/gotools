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
var daySeperator = NewString(".")
var timeSeperator = NewString(":")
var msSeperator = NewString(" ")

// represent a Duration of Time from Days to Milliseconds
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

// create a empty Time Span with zero values
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

// create a new TimeSpan from days, hours, minutes, seconds and milliseconds
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

// create a TimeSpan from a duration in Milliseconds
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

// create a new Time Span from the String "Day.Hour:Minute:Second Millisecond"
// values that not in the String was defined with 0
func NewTimeSpanFromString(isoString String) (*TimeSpan, error) {
	t, err := isoStringReader(isoString)
	if err != nil {
		return EmptyTimeSpan(), errors.New("Error on TimeSpan NewTimeSpanFromISOString: " + err.Error())
	}
	return t, nil
}

// is the Time Span a valid Time Span
func (ts *TimeSpan) IsValid() bool {
	return ts.isValid
}

// convert the TimeSpan to a String representation in the Format Day.Hour:Minute:Second Millisecond
func (ts *TimeSpan) AsString() *String {
	day := NewString(ts.day)
	hour := NewString(ts.hour).PadLeft(*NewNumber(2), *NewString("0"))
	minute := NewString(ts.minute).PadLeft(*NewNumber(2), *NewString("0"))
	second := NewString(ts.second).PadLeft(*NewNumber(2), *NewString("0"))
	millisecond := NewString(ts.millisecond).PadLeft(*NewNumber(3), *NewString("0"))
	res := EmptyString().
		Concat(*hour).Concat(*timeSeperator).
		Concat(*minute).Concat(*timeSeperator).
		Concat(*second)
	if millisecond.AsNumber().IsAbove(ZERO) {
		res = res.Concat(*msSeperator).Concat(*millisecond)
	}
	if day.AsNumber().IsAbove(ZERO) {
		res = EmptyString().Concat(*day).Concat(*daySeperator).Concat(*res)
	}
	return res
}

// get the number of Days in the TimeSpan
func (ts *TimeSpan) Day() *Number {
	return NewNumber(ts.day)
}

// get the number of Hours in the TimeSpan
func (ts *TimeSpan) Hour() *Number {
	return NewNumber(ts.hour)
}

// get the number of Minutes in the TimeSpan
func (ts *TimeSpan) Minute() *Number {
	return NewNumber(ts.minute)
}

// get the number of Seconds in the TimeSpan
func (ts *TimeSpan) Second() *Number {
	return NewNumber(ts.second)
}

// get the number of Milliseconds in the TimeSpan
func (ts *TimeSpan) Millisecond() *Number {
	return NewNumber(ts.millisecond)
}

// get the Sum of the Complete TimeSpan in Days
func (ts *TimeSpan) TotalDays() *Number {
	return ts.TotalMilliseconds().Divide(*millisecondsPerDay)
}

// get the Sum of the Complete TimeSpan in Hours
func (ts *TimeSpan) TotalHours() *Number {
	return ts.TotalMilliseconds().Divide(*millisecondsPerHour)
}

// get the Sum of the Complete TimeSpan in Minutes
func (ts *TimeSpan) TotalMinutes() *Number {
	return ts.TotalMilliseconds().Divide(*millisecondsPerMinute)
}

// get the Sum of the Complete TimeSpan in Seconds
func (ts *TimeSpan) TotalSeconds() *Number {
	return ts.TotalMilliseconds().Divide(*millisecondsPerSecond)
}

// get the Sum of the Complete TimeSpan in Milliseconds
func (ts *TimeSpan) TotalMilliseconds() *Number {
	return NewNumber(ts.day).Multiply(*millisecondsPerDay).Add(
		*NewNumber(ts.hour).Multiply(*millisecondsPerHour).Add(
			*NewNumber(ts.minute).Multiply(*millisecondsPerMinute).Add(
				*NewNumber(ts.second).Multiply(*millisecondsPerSecond).Add(
					*NewNumber(ts.millisecond)))))
}

// check is the TimeSpan Before the given TimeSpan
func (ts *TimeSpan) IsBefore(duration TimeSpan) bool {
	return ts.TotalMilliseconds().IsBelow(*duration.TotalMilliseconds())
}

// check is the TimeSpan After the given TimeSpan
func (ts *TimeSpan) IsAfter(duration TimeSpan) bool {
	return ts.TotalMilliseconds().IsAbove(*duration.TotalMilliseconds())
}

// check is the TimeSpan the given TimeSpan
func (ts *TimeSpan) Equals(duration TimeSpan) bool {
	return ts.TotalMilliseconds().Equals(*duration.TotalMilliseconds())
}

// add the given TimeSpan to the TimeSpan
func (ts *TimeSpan) Add(duration TimeSpan) *TimeSpan {
	return NewTimeSpanFromMilliseconds(*ts.TotalMilliseconds().Add(*duration.TotalMilliseconds()))
}

// remove the given TimeSpan to the TimeSpan
func (ts *TimeSpan) Subtract(duration TimeSpan) *TimeSpan {
	return NewTimeSpanFromMilliseconds(*ts.TotalMilliseconds().Subtract(*duration.TotalMilliseconds()))
}

// negate the TimeSpan day, hours, minutes, seconds and milliseconds
func (ts *TimeSpan) Negate() *TimeSpan {
	return NewTimeSpan(
		*NewNumber(-ts.day),
		*NewNumber(-ts.hour),
		*NewNumber(-ts.minute),
		*NewNumber(-ts.second),
		*NewNumber(-ts.millisecond))
}

func isoStringReader(value String) (*TimeSpan, error) {
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
