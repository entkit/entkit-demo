// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/entkit/entkit-demo/ent/company"
	"github.com/entkit/entkit-demo/ent/country"
	"github.com/entkit/entkit-demo/ent/phone"
	"github.com/entkit/entkit-demo/ent/predicate"
	"github.com/google/uuid"
)

// PhoneQuery is the builder for querying Phone entities.
type PhoneQuery struct {
	config
	ctx         *QueryContext
	order       []phone.Order
	inters      []Interceptor
	predicates  []predicate.Phone
	withCompany *CompanyQuery
	withCountry *CountryQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	loadTotal   []func(context.Context, []*Phone) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PhoneQuery builder.
func (pq *PhoneQuery) Where(ps ...predicate.Phone) *PhoneQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PhoneQuery) Limit(limit int) *PhoneQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PhoneQuery) Offset(offset int) *PhoneQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PhoneQuery) Unique(unique bool) *PhoneQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PhoneQuery) Order(o ...phone.Order) *PhoneQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryCompany chains the current query on the "company" edge.
func (pq *PhoneQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(phone.Table, phone.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, phone.CompanyTable, phone.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCountry chains the current query on the "country" edge.
func (pq *PhoneQuery) QueryCountry() *CountryQuery {
	query := (&CountryClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(phone.Table, phone.FieldID, selector),
			sqlgraph.To(country.Table, country.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, phone.CountryTable, phone.CountryColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Phone entity from the query.
// Returns a *NotFoundError when no Phone was found.
func (pq *PhoneQuery) First(ctx context.Context) (*Phone, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{phone.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PhoneQuery) FirstX(ctx context.Context) *Phone {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Phone ID from the query.
// Returns a *NotFoundError when no Phone ID was found.
func (pq *PhoneQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{phone.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PhoneQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Phone entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Phone entity is found.
// Returns a *NotFoundError when no Phone entities are found.
func (pq *PhoneQuery) Only(ctx context.Context) (*Phone, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{phone.Label}
	default:
		return nil, &NotSingularError{phone.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PhoneQuery) OnlyX(ctx context.Context) *Phone {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Phone ID in the query.
// Returns a *NotSingularError when more than one Phone ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PhoneQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{phone.Label}
	default:
		err = &NotSingularError{phone.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PhoneQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Phones.
func (pq *PhoneQuery) All(ctx context.Context) ([]*Phone, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Phone, *PhoneQuery]()
	return withInterceptors[[]*Phone](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PhoneQuery) AllX(ctx context.Context) []*Phone {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Phone IDs.
func (pq *PhoneQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(phone.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PhoneQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PhoneQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PhoneQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PhoneQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PhoneQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PhoneQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PhoneQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PhoneQuery) Clone() *PhoneQuery {
	if pq == nil {
		return nil
	}
	return &PhoneQuery{
		config:      pq.config,
		ctx:         pq.ctx.Clone(),
		order:       append([]phone.Order{}, pq.order...),
		inters:      append([]Interceptor{}, pq.inters...),
		predicates:  append([]predicate.Phone{}, pq.predicates...),
		withCompany: pq.withCompany.Clone(),
		withCountry: pq.withCountry.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PhoneQuery) WithCompany(opts ...func(*CompanyQuery)) *PhoneQuery {
	query := (&CompanyClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withCompany = query
	return pq
}

// WithCountry tells the query-builder to eager-load the nodes that are connected to
// the "country" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PhoneQuery) WithCountry(opts ...func(*CountryQuery)) *PhoneQuery {
	query := (&CountryClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withCountry = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Phone.Query().
//		GroupBy(phone.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PhoneQuery) GroupBy(field string, fields ...string) *PhoneGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PhoneGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = phone.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Phone.Query().
//		Select(phone.FieldTitle).
//		Scan(ctx, &v)
func (pq *PhoneQuery) Select(fields ...string) *PhoneSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PhoneSelect{PhoneQuery: pq}
	sbuild.label = phone.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PhoneSelect configured with the given aggregations.
func (pq *PhoneQuery) Aggregate(fns ...AggregateFunc) *PhoneSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PhoneQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !phone.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PhoneQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Phone, error) {
	var (
		nodes       = []*Phone{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withCompany != nil,
			pq.withCountry != nil,
		}
	)
	if pq.withCompany != nil || pq.withCountry != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, phone.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Phone).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Phone{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withCompany; query != nil {
		if err := pq.loadCompany(ctx, query, nodes, nil,
			func(n *Phone, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withCountry; query != nil {
		if err := pq.loadCountry(ctx, query, nodes, nil,
			func(n *Phone, e *Country) { n.Edges.Country = e }); err != nil {
			return nil, err
		}
	}
	for i := range pq.loadTotal {
		if err := pq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PhoneQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*Phone, init func(*Phone), assign func(*Phone, *Company)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Phone)
	for i := range nodes {
		if nodes[i].company_phones == nil {
			continue
		}
		fk := *nodes[i].company_phones
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(company.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "company_phones" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *PhoneQuery) loadCountry(ctx context.Context, query *CountryQuery, nodes []*Phone, init func(*Phone), assign func(*Phone, *Country)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Phone)
	for i := range nodes {
		if nodes[i].country_phones == nil {
			continue
		}
		fk := *nodes[i].country_phones
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(country.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "country_phones" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PhoneQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PhoneQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(phone.Table, phone.Columns, sqlgraph.NewFieldSpec(phone.FieldID, field.TypeUUID))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, phone.FieldID)
		for i := range fields {
			if fields[i] != phone.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PhoneQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(phone.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = phone.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PhoneGroupBy is the group-by builder for Phone entities.
type PhoneGroupBy struct {
	selector
	build *PhoneQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PhoneGroupBy) Aggregate(fns ...AggregateFunc) *PhoneGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PhoneGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PhoneQuery, *PhoneGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PhoneGroupBy) sqlScan(ctx context.Context, root *PhoneQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PhoneSelect is the builder for selecting fields of Phone entities.
type PhoneSelect struct {
	*PhoneQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PhoneSelect) Aggregate(fns ...AggregateFunc) *PhoneSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PhoneSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PhoneQuery, *PhoneSelect](ctx, ps.PhoneQuery, ps, ps.inters, v)
}

func (ps *PhoneSelect) sqlScan(ctx context.Context, root *PhoneQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
