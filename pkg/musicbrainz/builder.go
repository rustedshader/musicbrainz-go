package musicbrainzgo

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type QueryBuilder struct {
	params     []string
	searchType string
	format     string
	limit      int
	offset     int
}

func NewQueryBuilder(searchType string) *QueryBuilder {
	return &QueryBuilder{
		params:     make([]string, 0),
		searchType: searchType,
		format:     "json",
	}
}

func (qb *QueryBuilder) AddParam(field, value string) *QueryBuilder {
	if value != "" {
		qb.params = append(qb.params, fmt.Sprintf("%s:%s", field, value))
	}
	return qb
}
func (qb *QueryBuilder) AddNumericParam(field string, value int) *QueryBuilder {
	if value > 0 {
		qb.params = append(qb.params, fmt.Sprintf("%s:%d", field, value))
	}
	return qb
}

func (qb *QueryBuilder) AddRangeParam(field string, min, max int) *QueryBuilder {
	qb.params = append(qb.params, fmt.Sprintf("%s:[%d TO %d]", field, min, max))
	return qb
}

func (qb *QueryBuilder) SetLimit(limit int) *QueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *QueryBuilder) SetOffset(offset int) *QueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *QueryBuilder) Build() (string, error) {
	if qb.searchType == "" {
		return "", fmt.Errorf("search type cannot be empty")
	}
	u, err := url.Parse(API_ROOT_URL)
	if err != nil {
		return "", fmt.Errorf("Invalid API Base Url %w", err)
	}
	u.Path += fmt.Sprintf("/%s/", qb.searchType)

	q := url.Values{}

	if len(qb.params) > 0 {
		q.Set("query", strings.Join(qb.params, " AND "))
	}
	q.Set("fmt", qb.format)
	if qb.limit > 0 {
		q.Set("limit", strconv.Itoa(qb.limit))
	}
	if qb.offset > 0 {
		q.Set("offset", strconv.Itoa(qb.offset))
	}
	u.RawQuery = q.Encode()

	return u.String(), nil

}
