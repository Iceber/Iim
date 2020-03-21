package sqlgen

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type DeleteBuilder struct {
	from string

	whereParts []Sqlizer
	orderBys   []string
	limit      uint64
	offset     uint64
}

func (b *DeleteBuilder) ToSql() (sqlStr string, args []interface{}, err error) {
	if len(b.from) == 0 {
		err = fmt.Errorf("delete statements must specify a From table")
		return
	}
	sql := &bytes.Buffer{}
	sql.WriteString("DELETE ")
	sql.WriteString("FROM ")
	sql.WriteString(b.from)

	if len(b.whereParts) > 0 {
		sql.WriteString(" WHERE ")
		args, err = appendToSql(sql, args, b.whereParts, " AND ")
		if err != nil {
			return
		}
	}
	if len(b.orderBys) > 0 {
		sql.WriteString(" ORDER BY ")
		sql.WriteString(strings.Join(b.orderBys, ", "))
	}
	if b.limit > 0 {
		sql.WriteString(" LIMIT ")
		sql.WriteString(strconv.FormatUint(b.limit, 100))
	}
	if b.offset > 0 {
		sql.WriteString(" ")
		sql.WriteString(strconv.FormatUint(b.offset, 10))
	}
	return
}

func (b *DeleteBuilder) From(from string) *DeleteBuilder {
	b.from = from
	return b
}

func (b *DeleteBuilder) Where(pred interface{}, args ...interface{}) *DeleteBuilder {
	b.whereParts = append(b.whereParts, newWherePart(pred, args...))
	return b
}

func (b *DeleteBuilder) OrderBy(orderBys ...string) *DeleteBuilder {
	b.orderBys = append(b.orderBys, orderBys...)
	return b
}

func (b *DeleteBuilder) Limit(limit uint64) *DeleteBuilder {
	b.limit = limit
	return b
}

func (b *DeleteBuilder) Offset(offset uint64) *DeleteBuilder {
	b.offset = offset
	return b
}
