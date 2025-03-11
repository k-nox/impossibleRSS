package storage

import (
	"database/sql"
	"time"
)

func NewNullString(s *string) sql.NullString {
	if s == nil || *s == "" {
		return sql.NullString{}
	}

	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func NewNullTime(t *time.Time) sql.NullTime {
	if t == nil || t.IsZero() {
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  *t,
		Valid: true,
	}
}
