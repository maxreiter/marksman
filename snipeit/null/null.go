package null

import "encoding/json"

type NullableStringData struct {
	Val  string
	Init bool
}

type NullableString = *NullableStringData

var NullString = &NullableStringData{}

func NewNullableString(str string) NullableString {
	return &NullableStringData{
		Val:  str,
		Init: true,
	}
}

func (s NullableStringData) MarshalJSON() ([]byte, error) {
	if !s.Init {
		return []byte("null"), nil
	}

	return json.Marshal(s.Val)
}

func (s *NullableStringData) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*s = *NullString
		return nil
	}

	s.Init = true
	return json.Unmarshal(b, &s.Val)
}

type NullableIntData struct {
	Val  int
	Init bool
}

type NullableInt = *NullableIntData

var NullInt = &NullableIntData{}

func NewNullableInt(i int) NullableInt {
	return &NullableIntData{
		Val:  i,
		Init: true,
	}
}

func (i NullableIntData) MarshalJSON() ([]byte, error) {
	if !i.Init {
		return []byte("null"), nil
	}

	return json.Marshal(i.Val)
}

func (i *NullableIntData) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*i = *NullInt
		return nil
	}

	i.Init = true
	return json.Unmarshal(b, &i.Val)
}

type NullableInt32Data struct {
	Val  int32
	Init bool
}

type NullableInt32 = *NullableInt32Data

var NullInt32 = &NullableInt32Data{}

func NewNullableInt32(i int32) NullableInt32 {
	return &NullableInt32Data{
		Val:  i,
		Init: true,
	}
}

func (i NullableInt32Data) MarshalJSON() ([]byte, error) {
	if !i.Init {
		return []byte("null"), nil
	}

	return json.Marshal(i.Val)
}

func (i *NullableInt32Data) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*i = *NullInt32
		return nil
	}

	i.Init = true
	return json.Unmarshal(b, &i.Val)
}
