package types

import (
	"strings"
)

// Custom serializable bool type
type Bool struct {
	value bool
}

var (
	TrueBool  Bool
	FalseBool Bool
)

func init() {
	TrueBool = Bool{value: true}
	FalseBool = Bool{value: false}
}

func (b *Bool) MarshalJSON() ([]byte, error) {
	if b.value {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

func (b *Bool) UnmarshalJSON(data []byte) (err error) {
	s := strings.ToLower(strings.Trim(string(data), "\""))
	b.value = (s == "s" || s == "y" || s == "1" || s == "true" || s == "t")
	return nil
}

func (b *Bool) IsTrue() bool {
	return b.value
}

func FromNativeBool(value bool) Bool {
	return Bool{value: value}
}
