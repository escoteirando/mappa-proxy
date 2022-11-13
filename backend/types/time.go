package types

import (
	"fmt"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

var layouts = []string{
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	"Mon Jan _2 2006 15:04:05 MST",
}

/*
	"dataNascimento": "Sat Feb 05 1977 00:00:00 GMT+0000 (Coordinated Universal Time)",
	"dataValidade": "2022-01-01T00:00:00.000Z",
*/

// "Mon Jan _2 2006 15:04:05 MST"

// 	ANSIC       = "Mon Jan _2 15:04:05 2006"
// 	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
// 	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
// 	RFC822      = "02 Jan 06 15:04 MST"
// 	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
// 	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
// 	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
// 	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
// 	RFC3339     = "2006-01-02T15:04:05Z07:00"
// 	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
// 	Kitchen     = "3:04PM"
// 	// Handy time stamps.
// 	Stamp      = "Jan _2 15:04:05"
// 	StampMilli = "Jan _2 15:04:05.000"
// 	StampMicro = "Jan _2 15:04:05.000000"
// 	StampNano  = "Jan _2 15:04:05.000000000"

func Date(t time.Time) Time {
	return Time{Time: t}
}

func (dt *Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", dt.Format(time.RFC3339))), nil
}

func (dt *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" || len(s) == 0 {
		dt.Time = time.Time{}
		return
	}
	if len(s) > 33 {
		s = s[:33]
	}
	//Sat Feb 05 1977 00:00:00 GMT+0000
	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			*dt = Date(t)
			return nil
		}
	}
	dt.Time = time.Time{}
	return nil
}
