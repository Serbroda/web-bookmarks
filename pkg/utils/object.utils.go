package utils

import (
	"database/sql"
	"time"
)

func NullStringToString(val sql.NullString) *string {
	if val.Valid {
		return &val.String
	}
	return nil
}

func NullTimeToTime(val sql.NullTime) *time.Time {
	if val.Valid {
		return &val.Time
	}
	return nil
}

func NullBoolToBool(val sql.NullBool) *bool {
	if val.Valid {
		return &val.Bool
	}
	return nil
}

func NullInt16ToInt16(val sql.NullInt16) *int16 {
	if val.Valid {
		return &val.Int16
	}
	return nil
}

func NullInt32ToInt32(val sql.NullInt32) *int32 {
	if val.Valid {
		return &val.Int32
	}
	return nil
}

func NullInt64ToInt64(val sql.NullInt64) *int64 {
	if val.Valid {
		return &val.Int64
	}
	return nil
}
