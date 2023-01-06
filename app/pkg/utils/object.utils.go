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

func StringToNullString(val *string) sql.NullString {
	if val != nil {
		return sql.NullString{
			String: *val,
			Valid:  true,
		}
	}
	return sql.NullString{
		Valid: false,
	}
}

func TimeToNullTime(val *time.Time) sql.NullTime {
	if val != nil {
		return sql.NullTime{
			Time:  *val,
			Valid: true,
		}
	}
	return sql.NullTime{
		Valid: false,
	}
}

func BoolToNullBool(val *bool) sql.NullBool {
	if val != nil {
		return sql.NullBool{
			Bool:  *val,
			Valid: true,
		}
	}
	return sql.NullBool{
		Valid: false,
	}
}

func Int16ToNullInt16(val *int16) sql.NullInt16 {
	if val != nil {
		return sql.NullInt16{
			Int16: *val,
			Valid: true,
		}
	}
	return sql.NullInt16{
		Valid: false,
	}
}

func Int32ToNullInt32(val *int32) sql.NullInt32 {
	if val != nil {
		return sql.NullInt32{
			Int32: *val,
			Valid: true,
		}
	}
	return sql.NullInt32{
		Valid: false,
	}
}

func Int64ToNullInt64(val *int64) sql.NullInt64 {
	if val != nil {
		return sql.NullInt64{
			Int64: *val,
			Valid: true,
		}
	}
	return sql.NullInt64{
		Valid: false,
	}
}
