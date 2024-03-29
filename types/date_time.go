package types

import (
	"time"
)

var highMonths = []*Number{
	NewNumber(1),
	NewNumber(3),
	NewNumber(5),
	NewNumber(7),
	NewNumber(8),
	NewNumber(10),
	NewNumber(12),
}

type DateTime struct {
	day      int
	month    int
	year     int
	time     *TimeSpan
	location *time.Location
	isValid  bool
}

func EmptyDateTime() *DateTime {
	return &DateTime{
		day:      0,
		month:    0,
		year:     0,
		time:     EmptyTimeSpan(),
		location: time.UTC,
		isValid:  false,
	}
}

func NewDateTime(location String, year, month, day, hour, minute, second, millisecond Number) *DateTime {
	loc, err := time.LoadLocation(location.value)
	if err != nil {
		println("Error on DateTime NewDateTime: ", err.Error())
		return EmptyDateTime()
	}
	t := NewTimeSpan(*NewNumber(0), hour, minute, second, millisecond)
	return &DateTime{
		day:      day.AsInt(),
		month:    month.AsInt(),
		year:     year.AsInt(),
		time:     t,
		location: loc,
		isValid:  true,
	}
}

func (dt *DateTime) IsValid() bool {
	return dt.isValid
}

func (dt *DateTime) Day() *Number {
	return NewNumber(dt.day)
}

func (dt *DateTime) Month() *Number {
	return NewNumber(dt.month)
}

func (dt *DateTime) Year() *Number {
	return NewNumber(dt.year)
}

func (dt *DateTime) Hour() *Number {
	return dt.time.Hour()
}

func (dt *DateTime) Minute() *Number {
	return dt.time.Minute()
}

func (dt *DateTime) Second() *Number {
	return dt.time.Second()
}

func (dt *DateTime) Millisecond() *Number {
	return dt.time.Millisecond()
}

func (dt *DateTime) Clone() *DateTime {
	return NewDateTime(
		*NewString(dt.location.String()),
		*dt.Year(),
		*dt.Month(),
		*dt.Day(),
		*dt.Hour(),
		*dt.Minute(),
		*dt.Second(),
		*dt.Millisecond())
}

func (dt *DateTime) AddYears(years Number) {
	dt.SetYear(*dt.Year().Add(years))
}

func (dt *DateTime) AddMonths(months Number) {
	appendTime(dt.Month, months, *NewNumber(12), dt.AddYears, dt.SetMonth)
}

func (dt *DateTime) AddDays(days Number) {
	appendTime(dt.Day, days, *maxDayByMonth(dt.Month(), dt.Year()), dt.AddMonths, dt.SetDay)
}

func (dt *DateTime) AddHours(hours Number) {
	appendTime(dt.Hour, hours, *dt.time.HoursPerDay, dt.AddDays, dt.SetHour)
}

func (dt *DateTime) AddMinutes(minutes Number) {
	appendTime(dt.Minute, minutes, *dt.time.MinutesPerHour, dt.AddHours, dt.SetMinute)
}

func (dt *DateTime) AddSeconds(seconds Number) {
	appendTime(dt.Second, seconds, *dt.time.SecondsPerMinute, dt.AddMinutes, dt.SetSecond)
}

func (dt *DateTime) AddMilliseconds(milliseconds Number) {
	appendTime(dt.Millisecond, milliseconds, *dt.time.MillisecondsPerSecond, dt.AddSeconds, dt.SetMillisecond)
}

func (dt *DateTime) SetYear(year Number) {
	dt.year = year.AsInt()
}

func (dt *DateTime) SetMonth(month Number) {
	dt.month = month.AsInt()
}

func (dt *DateTime) SetDay(day Number) {
	dt.day = day.AsInt()
}

func (dt *DateTime) SetHour(hour Number) {
	dt.time.hour = hour.AsInt64()
}

func (dt *DateTime) SetMinute(minute Number) {
	dt.time.minute = minute.AsInt64()
}

func (dt *DateTime) SetSecond(second Number) {
	dt.time.second = second.AsInt64()
}

func (dt *DateTime) SetMillisecond(millisecond Number) {
	dt.time.millisecond = millisecond.AsInt64()
}

func (dt *DateTime) ToZone(location String) *DateTime {
	loc, err := time.LoadLocation(location.value)
	if err != nil {
		println("Error on DateTime NewDateTime: ", err.Error())
		return EmptyDateTime()
	}
	clone := dt.Clone()
	if dt.location != loc {
		clone.AddSeconds(*getUTCOffset(clone, loc))
		clone.location = loc
	}
	return clone
}

func (dt *DateTime) Equals(v DateTime) bool {
	return dt.Year().Equals(*v.Year()) &&
		dt.Month().Equals(*v.Month()) &&
		dt.Day().Equals(*v.Day()) &&
		dt.Hour().Equals(*v.Hour()) &&
		dt.Minute().Equals(*v.Minute()) &&
		dt.Second().Equals(*v.Second()) &&
		dt.Millisecond().Equals(*v.Millisecond())
}

func (dt *DateTime) AsUnixTimestamp() *Number {
	d := time.Date(
		dt.Year().AsInt(), time.Month(dt.Month().AsInt()), dt.Day().AsInt(),
		dt.Hour().AsInt(), dt.Minute().AsInt(), dt.Second().AsInt(), 0, dt.location)
	return NewNumber(d.UnixNano() / 1000000000)
}

func (dt *DateTime) AsString(fmt *String) *String {
	fmt = fmt.Replace(*NewString("YYYY"), *NewString(dt.Year()).
		PadLeft(*NewNumber(4), *NewString("0")))
	fmt = fmt.Replace(*NewString("YY"), *NewString(dt.Year()).
		SubString(*NewNumber(2), *NewNumber(2)).
		PadLeft(*NewNumber(2), *NewString("0")))
	fmt = fmt.Replace(*NewString("MM"), *NewString(dt.Month()).
		PadLeft(*NewNumber(2), *NewString("0")))
	fmt = fmt.Replace(*NewString("M"), *NewString(dt.Month()))
	fmt = fmt.Replace(*NewString("DD"), *NewString(dt.Day()).
		PadLeft(*NewNumber(2), *NewString("0")))
	fmt = fmt.Replace(*NewString("D"), *NewString(dt.Day()))
	fmt = fmt.Replace(*NewString("HH"), *NewString(dt.Hour()).
		PadLeft(*NewNumber(2), *NewString("0")))
	fmt = fmt.Replace(*NewString("H"), *NewString(dt.Hour()))
	fmt = fmt.Replace(*NewString("mm"), *NewString(dt.Minute()).
		PadLeft(*NewNumber(2), *NewString("0")))
	fmt = fmt.Replace(*NewString("m"), *NewString(dt.Minute()))
	fmt = fmt.Replace(*NewString("ss"), *NewString(dt.Second()).
		PadLeft(*NewNumber(2), *NewString("0")))
	fmt = fmt.Replace(*NewString("s"), *NewString(dt.Second()))
	fmt = fmt.Replace(*NewString("fff"), *NewString(dt.Millisecond()).
		PadLeft(*NewNumber(3), *NewString("0")))
	fmt = fmt.Replace(*NewString("ff"), *NewString(dt.Millisecond()).
		PadLeft(*NewNumber(2), *NewString("0")))
	fmt = fmt.Replace(*NewString("f"), *NewString(dt.Millisecond()))
	return fmt
}

func (dt *DateTime) IsBefore(v *DateTime) bool {
	return dt.AsUnixTimestamp().IsBelow(*v.AsUnixTimestamp())
}

func (dt *DateTime) IsAfter(v *DateTime) bool {
	return dt.AsUnixTimestamp().IsAbove(*v.AsUnixTimestamp())
}

func appendTime(getter func() *Number, value, border Number, addFn, setFn func(Number)) {
	toAdd := getter().Add(value)
	if toAdd.IsAbove(border) || toAdd.IsBelow(ZERO) {
		offset := toAdd.Divide(border).Floor(ZERO)
		remaind := toAdd.Subtract(*offset.Multiply(border))
		if remaind.Equals(ZERO) {
			remaind = NewNumber(1)
		}
		addFn(*offset)
		setFn(*remaind)
		return
	}
	setFn(*toAdd)
}

func isLeapYear(year *Number) bool {
	return year.Modulo(*NewNumber(4)).Equals(ZERO)
}

func maxDayByMonth(month, year *Number) *Number {
	isLeapYear := isLeapYear(year)
	isFebruary := month.Equals(*NewNumber(2))
	if isLeapYear && isFebruary {
		return NewNumber(29)
	} else if isFebruary {
		return NewNumber(28)
	}
	if month.IsIn(highMonths) {
		return NewNumber(31)
	}
	return NewNumber(30)
}

func getUTCOffset(dt *DateTime, target *time.Location) *Number {
	_, offset := time.Date(
		dt.Year().AsInt(),
		time.Month(dt.Month().AsInt()),
		dt.Day().AsInt(),
		dt.Hour().AsInt(),
		dt.Minute().AsInt(),
		dt.Second().AsInt(),
		0,
		dt.location,
	).Zone()
	_, offset2 := time.Date(
		dt.Year().AsInt(),
		time.Month(dt.Month().AsInt()),
		dt.Day().AsInt(),
		dt.Hour().AsInt(),
		dt.Minute().AsInt(),
		dt.Second().AsInt(),
		0,
		target,
	).Zone()
	if offset > 0 && offset2 == 0 {
		return NewNumber(-offset)
	}
	return NewNumber(offset2)
}
