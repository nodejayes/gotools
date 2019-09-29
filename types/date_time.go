package types

import "time"

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

func (dt *DateTime) AddYears(years Number) *DateTime {
	clone := dt.Clone()
	clone.SetYear(years)
	return clone
}

func (dt *DateTime) AddMonths(months Number) *DateTime {
	clone := dt.Clone()
	appendTime(months, *NewNumber(12), clone.SetMonth)
	return clone
}

func (dt *DateTime) AddDays(days Number) *DateTime {
	clone := dt.Clone()
	appendTime(days, *NewNumber(30), clone.SetDay)
	return clone
}

func (dt *DateTime) AddHours(hours Number) *DateTime {
	clone := dt.Clone()
	appendTime(hours, *dt.time.HoursPerDay, clone.SetHour)
	return clone
}

func (dt *DateTime) AddMinutes(minutes Number) *DateTime {
	clone := dt.Clone()
	appendTime(minutes, *dt.time.MinutesPerDay, clone.SetMinute)
	return clone
}

func (dt *DateTime) AddSeconds(seconds Number) *DateTime {
	clone := dt.Clone()
	appendTime(seconds, *dt.time.SecondsPerMinute, clone.SetSecond)
	return clone
}

func (dt *DateTime) AddMilliseconds(milliseconds Number) *DateTime {
	clone := dt.Clone()
	appendTime(milliseconds, *dt.time.MillisecondsPerSecond, clone.SetSecond)
	return clone
}

func (dt *DateTime) SetYear(year Number) *DateTime {
	clone := dt.Clone()
	clone.year = year.AsInt()
	return clone
}

func (dt *DateTime) SetMonth(month Number) *DateTime {
	clone := dt.Clone()
	clone.month = month.AsInt()
	return clone
}

func (dt *DateTime) SetDay(day Number) *DateTime {
	clone := dt.Clone()
	clone.day = day.AsInt()
	return clone
}

func (dt *DateTime) SetHour(hour Number) *DateTime {
	clone := dt.Clone()
	clone.time.hour = hour.AsInt64()
	return clone
}

func (dt *DateTime) SetMinute(minute Number) *DateTime {
	clone := dt.Clone()
	clone.time.minute = minute.AsInt64()
	return clone
}

func (dt *DateTime) SetSecond(second Number) *DateTime {
	clone := dt.Clone()
	clone.time.second = second.AsInt64()
	return clone
}

func (dt *DateTime) SetMillisecond(millisecond Number) *DateTime {
	clone := dt.Clone()
	clone.time.millisecond = millisecond.AsInt64()
	return clone
}

/*
func (dt *DateTime) InZone(location String) *DateTime {
	loc, err := time.LoadLocation(location.value)
	if err != nil {
		println("Error on DateTime NewDateTime: ", err.Error())
		return EmptyDateTime()
	}
	clone := dt.Clone()
	if dt.location != loc {
		clone.
	}
	return clone
}
*/

func appendTime(value, factor Number, setter func(Number) *DateTime) {
	border := factor.Subtract(*NewNumber(1))
	if value.IsAbove(*border) {
		offset := value.Divide(factor).Floor(ZERO)
		setter(*value.Subtract(*offset.Multiply(factor)))
	} else {
		setter(value)
	}
}
