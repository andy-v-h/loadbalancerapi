// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancer"
	"go.infratographer.com/x/idx"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[idx.PrefixedID]
	PageInfo       = entgql.PageInfo[idx.PrefixedID]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// LoadBalancerEdge is the edge representation of LoadBalancer.
type LoadBalancerEdge struct {
	Node   *LoadBalancer `json:"node"`
	Cursor Cursor        `json:"cursor"`
}

// LoadBalancerConnection is the connection containing edges to LoadBalancer.
type LoadBalancerConnection struct {
	Edges      []*LoadBalancerEdge `json:"edges"`
	PageInfo   PageInfo            `json:"pageInfo"`
	TotalCount int                 `json:"totalCount"`
}

func (c *LoadBalancerConnection) build(nodes []*LoadBalancer, pager *loadbalancerPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *LoadBalancer
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *LoadBalancer {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *LoadBalancer {
			return nodes[i]
		}
	}
	c.Edges = make([]*LoadBalancerEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &LoadBalancerEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// LoadBalancerPaginateOption enables pagination customization.
type LoadBalancerPaginateOption func(*loadbalancerPager) error

// WithLoadBalancerOrder configures pagination ordering.
func WithLoadBalancerOrder(order *LoadBalancerOrder) LoadBalancerPaginateOption {
	if order == nil {
		order = DefaultLoadBalancerOrder
	}
	o := *order
	return func(pager *loadbalancerPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultLoadBalancerOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithLoadBalancerFilter configures pagination filter.
func WithLoadBalancerFilter(filter func(*LoadBalancerQuery) (*LoadBalancerQuery, error)) LoadBalancerPaginateOption {
	return func(pager *loadbalancerPager) error {
		if filter == nil {
			return errors.New("LoadBalancerQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type loadbalancerPager struct {
	reverse bool
	order   *LoadBalancerOrder
	filter  func(*LoadBalancerQuery) (*LoadBalancerQuery, error)
}

func newLoadBalancerPager(opts []LoadBalancerPaginateOption, reverse bool) (*loadbalancerPager, error) {
	pager := &loadbalancerPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultLoadBalancerOrder
	}
	return pager, nil
}

func (p *loadbalancerPager) applyFilter(query *LoadBalancerQuery) (*LoadBalancerQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *loadbalancerPager) toCursor(lb *LoadBalancer) Cursor {
	return p.order.Field.toCursor(lb)
}

func (p *loadbalancerPager) applyCursors(query *LoadBalancerQuery, after, before *Cursor) (*LoadBalancerQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultLoadBalancerOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *loadbalancerPager) applyOrder(query *LoadBalancerQuery) *LoadBalancerQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultLoadBalancerOrder.Field {
		query = query.Order(DefaultLoadBalancerOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *loadbalancerPager) orderExpr(query *LoadBalancerQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultLoadBalancerOrder.Field {
			b.Comma().Ident(DefaultLoadBalancerOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to LoadBalancer.
func (lb *LoadBalancerQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...LoadBalancerPaginateOption,
) (*LoadBalancerConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newLoadBalancerPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if lb, err = pager.applyFilter(lb); err != nil {
		return nil, err
	}
	conn := &LoadBalancerConnection{Edges: []*LoadBalancerEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = lb.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if lb, err = pager.applyCursors(lb, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		lb.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := lb.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	lb = pager.applyOrder(lb)
	nodes, err := lb.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// LoadBalancerOrderFieldName orders LoadBalancer by name.
	LoadBalancerOrderFieldName = &LoadBalancerOrderField{
		Value: func(lb *LoadBalancer) (ent.Value, error) {
			return lb.Name, nil
		},
		column: loadbalancer.FieldName,
		toTerm: loadbalancer.ByName,
		toCursor: func(lb *LoadBalancer) Cursor {
			return Cursor{
				ID:    lb.ID,
				Value: lb.Name,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f LoadBalancerOrderField) String() string {
	var str string
	switch f.column {
	case LoadBalancerOrderFieldName.column:
		str = "NAME"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f LoadBalancerOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *LoadBalancerOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("LoadBalancerOrderField %T must be a string", v)
	}
	switch str {
	case "NAME":
		*f = *LoadBalancerOrderFieldName
	default:
		return fmt.Errorf("%s is not a valid LoadBalancerOrderField", str)
	}
	return nil
}

// LoadBalancerOrderField defines the ordering field of LoadBalancer.
type LoadBalancerOrderField struct {
	// Value extracts the ordering value from the given LoadBalancer.
	Value    func(*LoadBalancer) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) loadbalancer.OrderOption
	toCursor func(*LoadBalancer) Cursor
}

// LoadBalancerOrder defines the ordering of LoadBalancer.
type LoadBalancerOrder struct {
	Direction OrderDirection          `json:"direction"`
	Field     *LoadBalancerOrderField `json:"field"`
}

// DefaultLoadBalancerOrder is the default ordering of LoadBalancer.
var DefaultLoadBalancerOrder = &LoadBalancerOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &LoadBalancerOrderField{
		Value: func(lb *LoadBalancer) (ent.Value, error) {
			return lb.ID, nil
		},
		column: loadbalancer.FieldID,
		toTerm: loadbalancer.ByID,
		toCursor: func(lb *LoadBalancer) Cursor {
			return Cursor{ID: lb.ID}
		},
	},
}

// ToEdge converts LoadBalancer into LoadBalancerEdge.
func (lb *LoadBalancer) ToEdge(order *LoadBalancerOrder) *LoadBalancerEdge {
	if order == nil {
		order = DefaultLoadBalancerOrder
	}
	return &LoadBalancerEdge{
		Node:   lb,
		Cursor: order.Field.toCursor(lb),
	}
}
