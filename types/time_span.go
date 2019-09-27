package types

type TimeSpan struct {
	value                 int64
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

func NewTimeSpan(hours, minutes, seconds, milliseconds Number) *TimeSpan {

}

func NewTimeSpanFromMilliseconds(milliseconds Number) *TimeSpan {}
func NewTimeSpanFromISOString(isoString String) *TimeSpan       {}

func (ts *TimeSpan) IsValid() bool     {}
func (ts *TimeSpan) AsString() *String {}

func (ts *TimeSpan) Day() *Number               {}
func (ts *TimeSpan) TotalDays() *Number         {}
func (ts *TimeSpan) Hour() *Number              {}
func (ts *TimeSpan) TotalHours() *Number        {}
func (ts *TimeSpan) Minute() *Number            {}
func (ts *TimeSpan) TotalMinutes() *Number      {}
func (ts *TimeSpan) Second() *Number            {}
func (ts *TimeSpan) TotalSeconds() *Number      {}
func (ts *TimeSpan) Millisecond() *Number       {}
func (ts *TimeSpan) TotalMilliseconds() *Number {}

func (ts *TimeSpan) IsBefore(duration TimeSpan) bool      {}
func (ts *TimeSpan) IsAfter(duration TimeSpan) bool       {}
func (ts *TimeSpan) Negate() *TimeSpan                    {}
func (ts *TimeSpan) Add(duration TimeSpan) *TimeSpan      {}
func (ts *TimeSpan) Subtract(duration TimeSpan) *TimeSpan {}
func (ts *TimeSpan) Equals(duration TimeSpan) bool        {}
