package core

import (
	"database/sql"
	"encoding/json"
	"time"
)

// compose sql.NullString in NullString
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

// UnmarshalJSON for NullString
func (s *NullString) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *string
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		s.Valid = true
		s.String = *v
	} else {
		s.Valid = false
	}
	return nil
}

// compose sql.NullInt64 in NullInt64
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON for NullInt64
func (i NullInt64) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.Int64)
}

// UnmarshalJSON for NullInt64
func (i *NullInt64) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *int64
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		i.Valid = true
		i.Int64 = *v
	} else {
		i.Valid = false
	}
	return nil
}

// compose sql.NullFloat64 in NullInt64
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON for NullInt64
func (f NullFloat64) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(f.Float64)
}

// UnmarshalJSON for NullInt64
func (f *NullFloat64) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *float64
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		f.Valid = true
		f.Float64 = *v
	} else {
		f.Valid = false
	}
	return nil
}

// compose sql.NullTime in NullTime
type NullTime struct {
	sql.NullTime
}

// MarshalJSON for NullInt64
func (t NullTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time)
}

// UnmarshalJSON for NullInt64
func (t *NullTime) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *time.Time
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		t.Valid = true
		t.Time = *v
	} else {
		t.Valid = false
	}
	return nil
}
