package types

import (
	"errors"
	"strconv"
	"time"
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
	value                 time.Time
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
		value:                 time.Now().UTC(),
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
		value: time.Date(0,
			0,
			0,
			hours.AsInt(),
			minutes.AsInt(),
			seconds.AsInt(),
			millisecondsToNanoseconds(milliseconds).AsInt(),
			time.UTC),
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
	return &TimeSpan{
		value:                 time.Unix(0, millisecondsToNanoseconds(milliseconds).AsInt64()),
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

func NewTimeSpanFromISOString(isoString String) *TimeSpan {
	t, err := isoStringReader(isoString)
	if err != nil {
		println("Error on TimeSpan NewTimeSpanFromISOString: " + err.Error())
		return EmptyTimeSpan()
	}
	return &TimeSpan{
		value:                 t,
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

func (ts *TimeSpan) IsValid() bool {
	return ts.isValid
}

func (ts *TimeSpan) AsString() *String {
	println("missing Implementation")
	return EmptyString()
}

func (ts *TimeSpan) Day() *Number {
	return NewNumber(ts.value.Day())
}
func (ts *TimeSpan) Hour() *Number {
	return NewNumber(ts.value.Hour())
}
func (ts *TimeSpan) Minute() *Number {
	return NewNumber(ts.value.Minute())
}
func (ts *TimeSpan) Second() *Number {
	return NewNumber(ts.value.Second())
}
func (ts *TimeSpan) Millisecond() *Number {
	tmp := ts.value.Nanosecond() / int(time.Millisecond)
	return NewNumber(tmp)
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

func millisecondsToNanoseconds(milliseconds Number) *Number {
	nanoseconds := milliseconds.
		Multiply(*NewNumber(time.Millisecond)).
		Multiply(*NewNumber(time.Microsecond))
	return nanoseconds
}

func isoStringReader(value String) (time.Time, error) {
	daySeperator := NewString(".")
	timeSeperator := NewString(":")
	msSeperator := NewString(" ")

	isInLengthBorders := value.Length().IsAbove(*NewNumber(1)) && value.Length().IsBelow(*NewNumber(15))
	daySeperatorExists := value.ContainsCount(*daySeperator).Equals(*NewNumber(1))
	timeSeperatorExists := value.ContainsCount(*timeSeperator).Equals(*NewNumber(2))
	msSeperatorExists := value.ContainsCount(*msSeperator).Equals(*NewNumber(1))

	if !isInLengthBorders {
		println("Error on TimeSpan isoStringReader: string is to short or to long")
		return time.Now(), errors.New("string is to short or to long")
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

	return NewTimeSpan(), errors.New("missing Implementation")
}

/*
func isoStringWriter(value *TimeSpan) *String {
	tmp := EmptyString()
	if daySplit != nil && len(daySplit) > 1 {
		tmp = tmp.Concat(*daySplit[0].Concat(*daySeperator))
	} else {
		tmp = tmp.Concat(*NewString("0").Concat(*daySeperator))
	}
	if timeSplit != nil && len(timeSplit) > 2 {
		tmp = tmp.
			Concat(*timeSplit[0].Concat(*timeSeperator).
				Concat(*timeSplit[1].Concat(*timeSeperator).
					Concat(*timeSplit[2])))
	} else {
		tmp = tmp.
			Concat(*NewString("00").Concat(*timeSeperator).
				Concat(*NewString("00").Concat(*timeSeperator).
					Concat(*NewString("00"))))
	}
	if msSplit != nil && len(msSplit) > 1 {
		tmp = tmp.Concat(*msSeperator).Concat(*msSplit[1])
	}
	return tmp
}
*/
