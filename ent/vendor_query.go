// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/entkit/entkit-demo/ent/predicate"
	"github.com/entkit/entkit-demo/ent/product"
	"github.com/entkit/entkit-demo/ent/vendor"
	"github.com/entkit/entkit-demo/ent/warehouse"
	"github.com/google/uuid"
)

// VendorQuery is the builder for querying Vendor entities.
type VendorQuery struct {
	config
	ctx                 *QueryContext
	order               []vendor.Order
	inters              []Interceptor
	predicates          []predicate.Vendor
	withWarehouses      *WarehouseQuery
	withProducts        *ProductQuery
	modifiers           []func(*sql.Selector)
	loadTotal           []func(context.Context, []*Vendor) error
	withNamedWarehouses map[string]*WarehouseQuery
	withNamedProducts   map[string]*ProductQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VendorQuery builder.
func (vq *VendorQuery) Where(ps ...predicate.Vendor) *VendorQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VendorQuery) Limit(limit int) *VendorQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *VendorQuery) Offset(offset int) *VendorQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VendorQuery) Unique(unique bool) *VendorQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VendorQuery) Order(o ...vendor.Order) *VendorQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryWarehouses chains the current query on the "warehouses" edge.
func (vq *VendorQuery) QueryWarehouses() *WarehouseQuery {
	query := (&WarehouseClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vendor.Table, vendor.FieldID, selector),
			sqlgraph.To(warehouse.Table, warehouse.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, vendor.WarehousesTable, vendor.WarehousesColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProducts chains the current query on the "products" edge.
func (vq *VendorQuery) QueryProducts() *ProductQuery {
	query := (&ProductClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vendor.Table, vendor.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, vendor.ProductsTable, vendor.ProductsColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Vendor entity from the query.
// Returns a *NotFoundError when no Vendor was found.
func (vq *VendorQuery) First(ctx context.Context) (*Vendor, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vendor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VendorQuery) FirstX(ctx context.Context) *Vendor {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Vendor ID from the query.
// Returns a *NotFoundError when no Vendor ID was found.
func (vq *VendorQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vendor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VendorQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Vendor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Vendor entity is found.
// Returns a *NotFoundError when no Vendor entities are found.
func (vq *VendorQuery) Only(ctx context.Context) (*Vendor, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vendor.Label}
	default:
		return nil, &NotSingularError{vendor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VendorQuery) OnlyX(ctx context.Context) *Vendor {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Vendor ID in the query.
// Returns a *NotSingularError when more than one Vendor ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VendorQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vendor.Label}
	default:
		err = &NotSingularError{vendor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VendorQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Vendors.
func (vq *VendorQuery) All(ctx context.Context) ([]*Vendor, error) {
	ctx = setContextOp(ctx, vq.ctx, "All")
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Vendor, *VendorQuery]()
	return withInterceptors[[]*Vendor](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VendorQuery) AllX(ctx context.Context) []*Vendor {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Vendor IDs.
func (vq *VendorQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if vq.ctx.Unique == nil && vq.path != nil {
		vq.Unique(true)
	}
	ctx = setContextOp(ctx, vq.ctx, "IDs")
	if err = vq.Select(vendor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VendorQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VendorQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, "Count")
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VendorQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VendorQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VendorQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, "Exist")
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VendorQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VendorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VendorQuery) Clone() *VendorQuery {
	if vq == nil {
		return nil
	}
	return &VendorQuery{
		config:         vq.config,
		ctx:            vq.ctx.Clone(),
		order:          append([]vendor.Order{}, vq.order...),
		inters:         append([]Interceptor{}, vq.inters...),
		predicates:     append([]predicate.Vendor{}, vq.predicates...),
		withWarehouses: vq.withWarehouses.Clone(),
		withProducts:   vq.withProducts.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithWarehouses tells the query-builder to eager-load the nodes that are connected to
// the "warehouses" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VendorQuery) WithWarehouses(opts ...func(*WarehouseQuery)) *VendorQuery {
	query := (&WarehouseClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withWarehouses = query
	return vq
}

// WithProducts tells the query-builder to eager-load the nodes that are connected to
// the "products" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VendorQuery) WithProducts(opts ...func(*ProductQuery)) *VendorQuery {
	query := (&ProductClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withProducts = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Vendor.Query().
//		GroupBy(vendor.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vq *VendorQuery) GroupBy(field string, fields ...string) *VendorGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VendorGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = vendor.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Vendor.Query().
//		Select(vendor.FieldName).
//		Scan(ctx, &v)
func (vq *VendorQuery) Select(fields ...string) *VendorSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &VendorSelect{VendorQuery: vq}
	sbuild.label = vendor.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VendorSelect configured with the given aggregations.
func (vq *VendorQuery) Aggregate(fns ...AggregateFunc) *VendorSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VendorQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !vendor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VendorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Vendor, error) {
	var (
		nodes       = []*Vendor{}
		_spec       = vq.querySpec()
		loadedTypes = [2]bool{
			vq.withWarehouses != nil,
			vq.withProducts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Vendor).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Vendor{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withWarehouses; query != nil {
		if err := vq.loadWarehouses(ctx, query, nodes,
			func(n *Vendor) { n.Edges.Warehouses = []*Warehouse{} },
			func(n *Vendor, e *Warehouse) { n.Edges.Warehouses = append(n.Edges.Warehouses, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withProducts; query != nil {
		if err := vq.loadProducts(ctx, query, nodes,
			func(n *Vendor) { n.Edges.Products = []*Product{} },
			func(n *Vendor, e *Product) { n.Edges.Products = append(n.Edges.Products, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range vq.withNamedWarehouses {
		if err := vq.loadWarehouses(ctx, query, nodes,
			func(n *Vendor) { n.appendNamedWarehouses(name) },
			func(n *Vendor, e *Warehouse) { n.appendNamedWarehouses(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range vq.withNamedProducts {
		if err := vq.loadProducts(ctx, query, nodes,
			func(n *Vendor) { n.appendNamedProducts(name) },
			func(n *Vendor, e *Product) { n.appendNamedProducts(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range vq.loadTotal {
		if err := vq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VendorQuery) loadWarehouses(ctx context.Context, query *WarehouseQuery, nodes []*Vendor, init func(*Vendor), assign func(*Vendor, *Warehouse)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Vendor)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.InValues(vendor.WarehousesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.vendor_warehouses
		if fk == nil {
			return fmt.Errorf(`foreign-key "vendor_warehouses" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "vendor_warehouses" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vq *VendorQuery) loadProducts(ctx context.Context, query *ProductQuery, nodes []*Vendor, init func(*Vendor), assign func(*Vendor, *Product)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Vendor)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Product(func(s *sql.Selector) {
		s.Where(sql.InValues(vendor.ProductsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.vendor_products
		if fk == nil {
			return fmt.Errorf(`foreign-key "vendor_products" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "vendor_products" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (vq *VendorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VendorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(vendor.Table, vendor.Columns, sqlgraph.NewFieldSpec(vendor.FieldID, field.TypeUUID))
	_spec.From = vq.sql
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vq.path != nil {
		_spec.Unique = true
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vendor.FieldID)
		for i := range fields {
			if fields[i] != vendor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VendorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(vendor.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = vendor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedWarehouses tells the query-builder to eager-load the nodes that are connected to the "warehouses"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (vq *VendorQuery) WithNamedWarehouses(name string, opts ...func(*WarehouseQuery)) *VendorQuery {
	query := (&WarehouseClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if vq.withNamedWarehouses == nil {
		vq.withNamedWarehouses = make(map[string]*WarehouseQuery)
	}
	vq.withNamedWarehouses[name] = query
	return vq
}

// WithNamedProducts tells the query-builder to eager-load the nodes that are connected to the "products"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (vq *VendorQuery) WithNamedProducts(name string, opts ...func(*ProductQuery)) *VendorQuery {
	query := (&ProductClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if vq.withNamedProducts == nil {
		vq.withNamedProducts = make(map[string]*ProductQuery)
	}
	vq.withNamedProducts[name] = query
	return vq
}

// VendorGroupBy is the group-by builder for Vendor entities.
type VendorGroupBy struct {
	selector
	build *VendorQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VendorGroupBy) Aggregate(fns ...AggregateFunc) *VendorGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VendorGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, "GroupBy")
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VendorQuery, *VendorGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VendorGroupBy) sqlScan(ctx context.Context, root *VendorQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VendorSelect is the builder for selecting fields of Vendor entities.
type VendorSelect struct {
	*VendorQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VendorSelect) Aggregate(fns ...AggregateFunc) *VendorSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VendorSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, "Select")
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VendorQuery, *VendorSelect](ctx, vs.VendorQuery, vs, vs.inters, v)
}

func (vs *VendorSelect) sqlScan(ctx context.Context, root *VendorQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
