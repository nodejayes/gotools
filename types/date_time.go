package types

type DateTime struct {
	day     int
	month   int
	year    int
	time    *TimeSpan
	isValid bool
}

func NewDateTime(year, month, day, hour, minute, second, millisecond Number) *DateTime {
	t := NewTimeSpan(*NewNumber(0), hour, minute, second, millisecond)
	return &DateTime{
		day:     day.AsInt(),
		month:   month.AsInt(),
		year:    year.AsInt(),
		time:    t,
		isValid: true,
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
