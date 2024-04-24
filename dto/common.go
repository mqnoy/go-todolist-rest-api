package dto

import (
	"github.com/mqnoy/go-todolist-rest-api/model"
	"github.com/mqnoy/go-todolist-rest-api/util"
	"gopkg.in/guregu/null.v4"
)

type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"totalPages"`
	TotalItems int64 `json:"totalItems"`
	Offset     int   `json:",omitempty"`
}

type ListResponse[T any] struct {
	Rows     []*T       `json:"rows"`
	MetaData Pagination `json:"metaData"`
}

type SelectAndCount[M any] struct {
	Rows  []*M
	Count int64
}

type DetailParam struct {
	ID      string
	Session JwtPayload
}

// list param
type ListParam[T any] struct {
	Filters    T
	Orders     string
	Pagination Pagination
	Session    JwtPayload
}

type CreateParam[T any] struct {
	CreateValue T
	Session     JwtPayload
}

// update param
type UpdateParam[T any] struct {
	UID         string
	ID          string
	UpdateValue T
	Session     JwtPayload
}

type Timestamp struct {
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}

type FilterCommonParams struct {
	Keyword  string
	Name     string
	IsActive *bool
}

func ComposeTimestamp(m model.TimestampColumn) Timestamp {
	return Timestamp{
		CreatedAt: util.DateToEpoch(m.CreatedAt),
		UpdatedAt: util.DateToEpoch(m.UpdatedAt),
	}
}

func ParseNullTimeToEpoch(nullTime null.Time) null.Int {
	if nullTime.Valid {
		toUnix := util.DateToEpoch(nullTime.Time)
		return null.IntFrom(toUnix)
	}

	return null.Int{}
}
