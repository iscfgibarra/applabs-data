package repositories

import (
	"database/sql"
	"time"
)

func StringTimeToNull(s string) sql.NullString {
	layout := "2006-01-02"
	var t time.Time
	var err error
	var null sql.NullString

	t, err = time.Parse(layout, s)
	if err == nil {
		nullT := sql.NullTime{Time: t}
		null = sql.NullString{String: nullT.Time.Format("2006-01-02")}

		if !nullT.Time.IsZero() {
			null.Valid = true
		}
	}

	return null
}

func StringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}

	if null.String != "" {
		null.Valid = true
	}

	return null
}

func StringToTime(s string) time.Time {
	layout := time.RFC3339
	var t time.Time
	var err error

	t, err = time.Parse(layout, s)
	if err != nil {
		return time.Time{}
	}

	return t
}
