// Code generated by greenstar scripts; DO NOT EDIT.

package transaction

import (
	"errors"
	"fmt"
	"github.com/arikkfir/greenstar/backend/internal/server/util"
	"github.com/arikkfir/greenstar/backend/internal/util/lang"
	"github.com/arikkfir/greenstar/backend/internal/util/observability"
	"github.com/shopspring/decimal"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"
)

var (
	_ = decimal.Decimal{}
	_ = time.Time{}
)

var (
	sortableColumns = []string{"amount", "convertedAmount", "currency", "date", "description", "referenceId"}
)

type ListRequest struct {
	properties      []string
	Offset          *uint            `url:"_offset,omitempty"`
	Count           *uint            `url:"_count,omitempty"`
	Sort            []string         `url:"_sort,omitempty"`
	Description     *string          `url:"description,omitempty"`
	MaxAmount       *decimal.Decimal `url:"maxAmount,omitempty"`
	MaxDate         *time.Time       `url:"maxDate,omitempty"`
	MinAmount       *decimal.Decimal `url:"minAmount,omitempty"`
	MinDate         *time.Time       `url:"minDate,omitempty"`
	ReferenceID     *string          `url:"referenceId,omitempty"`
	SourceAccountID *string          `url:"sourceAccountId,omitempty"`
	TargetAccountID *string          `url:"targetAccountId,omitempty"`
	Currency        string           `url:"currency,omitempty"`
}

func (lr *ListRequest) HasDescription() bool { return slices.Contains(lr.properties, "description") }
func (lr *ListRequest) HasMaxAmount() bool   { return slices.Contains(lr.properties, "maxAmount") }
func (lr *ListRequest) HasMaxDate() bool     { return slices.Contains(lr.properties, "maxDate") }
func (lr *ListRequest) HasMinAmount() bool   { return slices.Contains(lr.properties, "minAmount") }
func (lr *ListRequest) HasMinDate() bool     { return slices.Contains(lr.properties, "minDate") }
func (lr *ListRequest) HasReferenceID() bool { return slices.Contains(lr.properties, "referenceId") }
func (lr *ListRequest) HasSourceAccountID() bool {
	return slices.Contains(lr.properties, "sourceAccountId")
}
func (lr *ListRequest) HasTargetAccountID() bool {
	return slices.Contains(lr.properties, "targetAccountId")
}
func (lr *ListRequest) UnmarshalFromRequest(r *http.Request) error {
	lr.properties = nil

	values := r.Form
	if values.Has("description") {
		lr.properties = append(lr.properties, "description")
		if rawValue := values.Get("description"); rawValue == util.QueryNilValue {
			lr.Description = nil
		} else {
			sv := lang.PtrOf(rawValue)
			lr.Description = sv
		}
	}
	if values.Has("maxAmount") {
		lr.properties = append(lr.properties, "maxAmount")
		if rawValue := values.Get("maxAmount"); rawValue == util.QueryNilValue {
			lr.MaxAmount = nil
		} else {
			sv := lang.PtrOf(rawValue)
			if dv, err := decimal.NewFromString(*sv); err != nil {
				return err
			} else {
				lr.MaxAmount = lang.PtrOf(dv)
			}
		}
	}
	if values.Has("maxDate") {
		lr.properties = append(lr.properties, "maxDate")
		if rawValue := values.Get("maxDate"); rawValue == util.QueryNilValue {
			lr.MaxDate = nil
		} else {
			sv := lang.PtrOf(rawValue)
			if tv, err := time.Parse(time.RFC3339, *sv); err != nil {
				return err
			} else {
				lr.MaxDate = lang.PtrOf(tv)
			}
		}
	}
	if values.Has("minAmount") {
		lr.properties = append(lr.properties, "minAmount")
		if rawValue := values.Get("minAmount"); rawValue == util.QueryNilValue {
			lr.MinAmount = nil
		} else {
			sv := lang.PtrOf(rawValue)
			if dv, err := decimal.NewFromString(*sv); err != nil {
				return err
			} else {
				lr.MinAmount = lang.PtrOf(dv)
			}
		}
	}
	if values.Has("minDate") {
		lr.properties = append(lr.properties, "minDate")
		if rawValue := values.Get("minDate"); rawValue == util.QueryNilValue {
			lr.MinDate = nil
		} else {
			sv := lang.PtrOf(rawValue)
			if tv, err := time.Parse(time.RFC3339, *sv); err != nil {
				return err
			} else {
				lr.MinDate = lang.PtrOf(tv)
			}
		}
	}
	if values.Has("referenceId") {
		lr.properties = append(lr.properties, "referenceId")
		if rawValue := values.Get("referenceId"); rawValue == util.QueryNilValue {
			lr.ReferenceID = nil
		} else {
			sv := lang.PtrOf(rawValue)
			lr.ReferenceID = sv
		}
	}
	if values.Has("sourceAccountId") {
		lr.properties = append(lr.properties, "sourceAccountId")
		if rawValue := values.Get("sourceAccountId"); rawValue == util.QueryNilValue {
			lr.SourceAccountID = nil
		} else {
			sv := lang.PtrOf(rawValue)
			lr.SourceAccountID = sv
		}
	}
	if values.Has("targetAccountId") {
		lr.properties = append(lr.properties, "targetAccountId")
		if rawValue := values.Get("targetAccountId"); rawValue == util.QueryNilValue {
			lr.TargetAccountID = nil
		} else {
			sv := lang.PtrOf(rawValue)
			lr.TargetAccountID = sv
		}
	}
	if values.Has("currency") {
		lr.properties = append(lr.properties, "currency")
		if rawValue := values.Get("currency"); rawValue == util.QueryNilValue {
			return fmt.Errorf("%w: '%s' is required", util.ErrBadRequest, "currency")
		} else {
			sv := rawValue
			lr.Currency = sv
		}
	} else {
		return fmt.Errorf("%w: '%s' is required", util.ErrBadRequest, "currency")
	}

	lr.Offset = nil
	if values.Has("_offset") {
		if v, err := strconv.ParseUint(values.Get("_offset"), 0, 0); err != nil {
			return fmt.Errorf("%w: invalid offset '%s'", util.ErrBadRequest, values.Get("_offset"))
		} else {
			lr.Offset = lang.PtrOf(uint(v))
		}
	}

	lr.Count = nil
	if values.Has("_count") {
		if v, err := strconv.ParseUint(values.Get("_count"), 0, 0); err != nil {
			return fmt.Errorf("%w: invalid count '%s'", util.ErrBadRequest, values.Get("_count"))
		} else {
			lr.Count = lang.PtrOf(uint(v))
		}
	}

	lr.Sort = nil
	if values.Has("_sort") {
		if sort := values.Get("_sort"); sort != "" {
			lr.Sort = strings.Split(sort, ",")
			for _, s := range lr.Sort {
				sortTokens := strings.Split(s, ":")
				var col, dir string
				if len(sortTokens) == 1 {
					col = sortTokens[0]
					dir = "asc"
				} else if len(sortTokens) == 2 {
					col = sortTokens[0]
					dir = strings.ToLower(sortTokens[1])
				} else {
					return fmt.Errorf("%w: invalid sort spec '%s'", util.ErrBadRequest, s)
				}
				if dir != "asc" && dir != "desc" {
					return fmt.Errorf("%w: invalid sort direction in '%s'", util.ErrBadRequest, s)
				} else if !slices.Contains(sortableColumns, col) {
					return fmt.Errorf("%w: column '%s' is not sortable", util.ErrBadRequest, col)
				}
			}
		}
	}

	return nil
}

type ListResponse struct {
	TotalCount uint          `json:"totalCount"`
	Items      []Transaction `json:"items"`
}

func (s *Server) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	l := observability.GetLogger(ctx)
	if err := util.VerifyPermissions(ctx, "transactions:read"); err != nil {
		util.ServeError(w, r, err)
		return
	}

	if err := r.ParseForm(); err != nil {
		util.ServeError(w, r, errors.Join(util.ErrBadRequest, err))
		return
	}

	req := ListRequest{}
	if err := req.UnmarshalFromRequest(r); err != nil {
		util.ServeError(w, r, err)
		return
	}

	res, err := s.h.List(ctx, req)
	if err != nil {
		if code := util.ServeError(w, r, err); code >= http.StatusInternalServerError {
			l.ErrorContext(ctx, "Failed listing transactions", "err", err)
		}
		return
	}

	type ResponseWithOffset struct {
		Offset *uint `json:"offset,omitempty"`
		ListResponse
	}

	resWithOffset := ResponseWithOffset{
		Offset:       req.Offset,
		ListResponse: *res,
	}
	if err := util.Marshal(w, r, http.StatusOK, resWithOffset); err != nil {
		l.ErrorContext(ctx, "Failed marshaling transactions", "err", err)
		util.ServeError(w, r, err)
	}
}
