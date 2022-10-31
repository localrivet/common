package types

import (
	"database/sql"
)

func NewNullString(value interface{}) sql.NullString {
	n := sql.NullString{}
	switch t := value.(type) {
	case string:
		n.String = t
		n.Valid = true
	case *string:
		n.Valid = t != nil
		if n.Valid {
			n.String = *t
		}
	}
	return n
}

func NewStringFromNull(value sql.NullString) string {
	if value.Valid {
		return value.String
	}
	return ""
}

func NewStringPtr(value string) *string {
	return &value
}

func NewNullInt64(value interface{}) sql.NullInt64 {
	n := sql.NullInt64{}
	switch t := value.(type) {
	case int64:
		n.Int64 = t
		n.Valid = true
	case *int64:
		n.Valid = t != nil
		if n.Valid {
			n.Int64 = *t
		}
	}
	return n
}

func NewInt64FromNull(value sql.NullInt64) int64 {
	if value.Valid {
		return value.Int64
	}
	return 0
}

func NewInt64Ptr(value int64) *int64 {
	return &value
}

func NewNullInt32(value interface{}) sql.NullInt32 {
	n := sql.NullInt32{}
	switch t := value.(type) {
	case int32:
		n.Int32 = t
		n.Valid = true
	case *int32:
		n.Valid = t != nil
		if n.Valid {
			n.Int32 = *t
		}
	}
	return n
}

func NewInt32FromNull(value sql.NullInt32) int32 {
	if value.Valid {
		return value.Int32
	}
	return 0
}

func NewInt32Ptr(value int32) *int32 {
	return &value
}

func NewNullBool(value interface{}) sql.NullBool {
	n := sql.NullBool{}
	switch t := value.(type) {
	case bool:
		n.Bool = t
		n.Valid = true
	case *bool:
		n.Valid = t != nil
		if n.Valid {
			n.Bool = *t
		}
	}
	return n
}

func NewBoolFromNull(value sql.NullBool) bool {
	if value.Valid {
		return value.Bool
	}
	return false
}

func NewBoolPtr(value bool) *bool {
	return &value
}
