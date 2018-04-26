package opentok

import (
	"strconv"
	"time"
)

// Timestamp ...
type Timestamp struct {
	time.Time
}

// UnmarshalJSON ...
func (t *Timestamp) UnmarshalJSON(b []byte) (err error) {
	str := string(b)

	if str == "null" {
		(*t).Time = time.Time{}
		return nil
	}

	var ustr string

	ustr, err = strconv.Unquote(str)
	if err != nil {
		ustr = str
	}

	m, err := time.Parse(time.RFC1123Z, ustr)

	if err == nil {
		(*t).Time = m
	} else {
		(*t).Time = time.Time{}
	}

	return nil
}

// IsZero ...
func (t *Timestamp) IsZero() bool {
	return t.Time.IsZero()
}

// Equal ...
func (t Timestamp) Equal(m Timestamp) bool {
	return t.Time.Equal(m.Time)
}
