// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/entkit/entkit-demo/ent/predicate"
	"github.com/entkit/entkit-demo/ent/product"
	"github.com/entkit/entkit-demo/ent/vendor"
	"github.com/entkit/entkit-demo/ent/warehouse"
	"github.com/google/uuid"
)

// WarehouseUpdate is the builder for updating Warehouse entities.
type WarehouseUpdate struct {
	config
	hooks    []Hook
	mutation *WarehouseMutation
}

// Where appends a list predicates to the WarehouseUpdate builder.
func (wu *WarehouseUpdate) Where(ps ...predicate.Warehouse) *WarehouseUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetName sets the "name" field.
func (wu *WarehouseUpdate) SetName(s string) *WarehouseUpdate {
	wu.mutation.SetName(s)
	return wu
}

// SetLastUpdate sets the "last_update" field.
func (wu *WarehouseUpdate) SetLastUpdate(t time.Time) *WarehouseUpdate {
	wu.mutation.SetLastUpdate(t)
	return wu
}

// SetNillableLastUpdate sets the "last_update" field if the given value is not nil.
func (wu *WarehouseUpdate) SetNillableLastUpdate(t *time.Time) *WarehouseUpdate {
	if t != nil {
		wu.SetLastUpdate(*t)
	}
	return wu
}

// ClearLastUpdate clears the value of the "last_update" field.
func (wu *WarehouseUpdate) ClearLastUpdate() *WarehouseUpdate {
	wu.mutation.ClearLastUpdate()
	return wu
}

// SetOriginalData sets the "original_data" field.
func (wu *WarehouseUpdate) SetOriginalData(s string) *WarehouseUpdate {
	wu.mutation.SetOriginalData(s)
	return wu
}

// SetNillableOriginalData sets the "original_data" field if the given value is not nil.
func (wu *WarehouseUpdate) SetNillableOriginalData(s *string) *WarehouseUpdate {
	if s != nil {
		wu.SetOriginalData(*s)
	}
	return wu
}

// ClearOriginalData clears the value of the "original_data" field.
func (wu *WarehouseUpdate) ClearOriginalData() *WarehouseUpdate {
	wu.mutation.ClearOriginalData()
	return wu
}

// SetEnabled sets the "enabled" field.
func (wu *WarehouseUpdate) SetEnabled(b bool) *WarehouseUpdate {
	wu.mutation.SetEnabled(b)
	return wu
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (wu *WarehouseUpdate) SetNillableEnabled(b *bool) *WarehouseUpdate {
	if b != nil {
		wu.SetEnabled(*b)
	}
	return wu
}

// SetFilters sets the "filters" field.
func (wu *WarehouseUpdate) SetFilters(s []string) *WarehouseUpdate {
	wu.mutation.SetFilters(s)
	return wu
}

// AppendFilters appends s to the "filters" field.
func (wu *WarehouseUpdate) AppendFilters(s []string) *WarehouseUpdate {
	wu.mutation.AppendFilters(s)
	return wu
}

// ClearFilters clears the value of the "filters" field.
func (wu *WarehouseUpdate) ClearFilters() *WarehouseUpdate {
	wu.mutation.ClearFilters()
	return wu
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (wu *WarehouseUpdate) AddProductIDs(ids ...uuid.UUID) *WarehouseUpdate {
	wu.mutation.AddProductIDs(ids...)
	return wu
}

// AddProducts adds the "products" edges to the Product entity.
func (wu *WarehouseUpdate) AddProducts(p ...*Product) *WarehouseUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wu.AddProductIDs(ids...)
}

// SetVendorID sets the "vendor" edge to the Vendor entity by ID.
func (wu *WarehouseUpdate) SetVendorID(id uuid.UUID) *WarehouseUpdate {
	wu.mutation.SetVendorID(id)
	return wu
}

// SetNillableVendorID sets the "vendor" edge to the Vendor entity by ID if the given value is not nil.
func (wu *WarehouseUpdate) SetNillableVendorID(id *uuid.UUID) *WarehouseUpdate {
	if id != nil {
		wu = wu.SetVendorID(*id)
	}
	return wu
}

// SetVendor sets the "vendor" edge to the Vendor entity.
func (wu *WarehouseUpdate) SetVendor(v *Vendor) *WarehouseUpdate {
	return wu.SetVendorID(v.ID)
}

// Mutation returns the WarehouseMutation object of the builder.
func (wu *WarehouseUpdate) Mutation() *WarehouseMutation {
	return wu.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (wu *WarehouseUpdate) ClearProducts() *WarehouseUpdate {
	wu.mutation.ClearProducts()
	return wu
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (wu *WarehouseUpdate) RemoveProductIDs(ids ...uuid.UUID) *WarehouseUpdate {
	wu.mutation.RemoveProductIDs(ids...)
	return wu
}

// RemoveProducts removes "products" edges to Product entities.
func (wu *WarehouseUpdate) RemoveProducts(p ...*Product) *WarehouseUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wu.RemoveProductIDs(ids...)
}

// ClearVendor clears the "vendor" edge to the Vendor entity.
func (wu *WarehouseUpdate) ClearVendor() *WarehouseUpdate {
	wu.mutation.ClearVendor()
	return wu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WarehouseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, WarehouseMutation](ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WarehouseUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WarehouseUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WarehouseUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WarehouseUpdate) check() error {
	if v, ok := wu.mutation.Name(); ok {
		if err := warehouse.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Warehouse.name": %w`, err)}
		}
	}
	return nil
}

func (wu *WarehouseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(warehouse.Table, warehouse.Columns, sqlgraph.NewFieldSpec(warehouse.FieldID, field.TypeUUID))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Name(); ok {
		_spec.SetField(warehouse.FieldName, field.TypeString, value)
	}
	if value, ok := wu.mutation.LastUpdate(); ok {
		_spec.SetField(warehouse.FieldLastUpdate, field.TypeTime, value)
	}
	if wu.mutation.LastUpdateCleared() {
		_spec.ClearField(warehouse.FieldLastUpdate, field.TypeTime)
	}
	if value, ok := wu.mutation.OriginalData(); ok {
		_spec.SetField(warehouse.FieldOriginalData, field.TypeString, value)
	}
	if wu.mutation.OriginalDataCleared() {
		_spec.ClearField(warehouse.FieldOriginalData, field.TypeString)
	}
	if value, ok := wu.mutation.Enabled(); ok {
		_spec.SetField(warehouse.FieldEnabled, field.TypeBool, value)
	}
	if value, ok := wu.mutation.Filters(); ok {
		_spec.SetField(warehouse.FieldFilters, field.TypeJSON, value)
	}
	if value, ok := wu.mutation.AppendedFilters(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, warehouse.FieldFilters, value)
		})
	}
	if wu.mutation.FiltersCleared() {
		_spec.ClearField(warehouse.FieldFilters, field.TypeJSON)
	}
	if wu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   warehouse.ProductsTable,
			Columns: []string{warehouse.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedProductsIDs(); len(nodes) > 0 && !wu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   warehouse.ProductsTable,
			Columns: []string{warehouse.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   warehouse.ProductsTable,
			Columns: []string{warehouse.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.VendorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   warehouse.VendorTable,
			Columns: []string{warehouse.VendorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vendor.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.VendorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   warehouse.VendorTable,
			Columns: []string{warehouse.VendorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vendor.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{warehouse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WarehouseUpdateOne is the builder for updating a single Warehouse entity.
type WarehouseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WarehouseMutation
}

// SetName sets the "name" field.
func (wuo *WarehouseUpdateOne) SetName(s string) *WarehouseUpdateOne {
	wuo.mutation.SetName(s)
	return wuo
}

// SetLastUpdate sets the "last_update" field.
func (wuo *WarehouseUpdateOne) SetLastUpdate(t time.Time) *WarehouseUpdateOne {
	wuo.mutation.SetLastUpdate(t)
	return wuo
}

// SetNillableLastUpdate sets the "last_update" field if the given value is not nil.
func (wuo *WarehouseUpdateOne) SetNillableLastUpdate(t *time.Time) *WarehouseUpdateOne {
	if t != nil {
		wuo.SetLastUpdate(*t)
	}
	return wuo
}

// ClearLastUpdate clears the value of the "last_update" field.
func (wuo *WarehouseUpdateOne) ClearLastUpdate() *WarehouseUpdateOne {
	wuo.mutation.ClearLastUpdate()
	return wuo
}

// SetOriginalData sets the "original_data" field.
func (wuo *WarehouseUpdateOne) SetOriginalData(s string) *WarehouseUpdateOne {
	wuo.mutation.SetOriginalData(s)
	return wuo
}

// SetNillableOriginalData sets the "original_data" field if the given value is not nil.
func (wuo *WarehouseUpdateOne) SetNillableOriginalData(s *string) *WarehouseUpdateOne {
	if s != nil {
		wuo.SetOriginalData(*s)
	}
	return wuo
}

// ClearOriginalData clears the value of the "original_data" field.
func (wuo *WarehouseUpdateOne) ClearOriginalData() *WarehouseUpdateOne {
	wuo.mutation.ClearOriginalData()
	return wuo
}

// SetEnabled sets the "enabled" field.
func (wuo *WarehouseUpdateOne) SetEnabled(b bool) *WarehouseUpdateOne {
	wuo.mutation.SetEnabled(b)
	return wuo
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (wuo *WarehouseUpdateOne) SetNillableEnabled(b *bool) *WarehouseUpdateOne {
	if b != nil {
		wuo.SetEnabled(*b)
	}
	return wuo
}

// SetFilters sets the "filters" field.
func (wuo *WarehouseUpdateOne) SetFilters(s []string) *WarehouseUpdateOne {
	wuo.mutation.SetFilters(s)
	return wuo
}

// AppendFilters appends s to the "filters" field.
func (wuo *WarehouseUpdateOne) AppendFilters(s []string) *WarehouseUpdateOne {
	wuo.mutation.AppendFilters(s)
	return wuo
}

// ClearFilters clears the value of the "filters" field.
func (wuo *WarehouseUpdateOne) ClearFilters() *WarehouseUpdateOne {
	wuo.mutation.ClearFilters()
	return wuo
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (wuo *WarehouseUpdateOne) AddProductIDs(ids ...uuid.UUID) *WarehouseUpdateOne {
	wuo.mutation.AddProductIDs(ids...)
	return wuo
}

// AddProducts adds the "products" edges to the Product entity.
func (wuo *WarehouseUpdateOne) AddProducts(p ...*Product) *WarehouseUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wuo.AddProductIDs(ids...)
}

// SetVendorID sets the "vendor" edge to the Vendor entity by ID.
func (wuo *WarehouseUpdateOne) SetVendorID(id uuid.UUID) *WarehouseUpdateOne {
	wuo.mutation.SetVendorID(id)
	return wuo
}

// SetNillableVendorID sets the "vendor" edge to the Vendor entity by ID if the given value is not nil.
func (wuo *WarehouseUpdateOne) SetNillableVendorID(id *uuid.UUID) *WarehouseUpdateOne {
	if id != nil {
		wuo = wuo.SetVendorID(*id)
	}
	return wuo
}

// SetVendor sets the "vendor" edge to the Vendor entity.
func (wuo *WarehouseUpdateOne) SetVendor(v *Vendor) *WarehouseUpdateOne {
	return wuo.SetVendorID(v.ID)
}

// Mutation returns the WarehouseMutation object of the builder.
func (wuo *WarehouseUpdateOne) Mutation() *WarehouseMutation {
	return wuo.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (wuo *WarehouseUpdateOne) ClearProducts() *WarehouseUpdateOne {
	wuo.mutation.ClearProducts()
	return wuo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (wuo *WarehouseUpdateOne) RemoveProductIDs(ids ...uuid.UUID) *WarehouseUpdateOne {
	wuo.mutation.RemoveProductIDs(ids...)
	return wuo
}

// RemoveProducts removes "products" edges to Product entities.
func (wuo *WarehouseUpdateOne) RemoveProducts(p ...*Product) *WarehouseUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wuo.RemoveProductIDs(ids...)
}

// ClearVendor clears the "vendor" edge to the Vendor entity.
func (wuo *WarehouseUpdateOne) ClearVendor() *WarehouseUpdateOne {
	wuo.mutation.ClearVendor()
	return wuo
}

// Where appends a list predicates to the WarehouseUpdate builder.
func (wuo *WarehouseUpdateOne) Where(ps ...predicate.Warehouse) *WarehouseUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WarehouseUpdateOne) Select(field string, fields ...string) *WarehouseUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Warehouse entity.
func (wuo *WarehouseUpdateOne) Save(ctx context.Context) (*Warehouse, error) {
	return withHooks[*Warehouse, WarehouseMutation](ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WarehouseUpdateOne) SaveX(ctx context.Context) *Warehouse {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WarehouseUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WarehouseUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WarehouseUpdateOne) check() error {
	if v, ok := wuo.mutation.Name(); ok {
		if err := warehouse.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Warehouse.name": %w`, err)}
		}
	}
	return nil
}

func (wuo *WarehouseUpdateOne) sqlSave(ctx context.Context) (_node *Warehouse, err error) {
	if err := wuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(warehouse.Table, warehouse.Columns, sqlgraph.NewFieldSpec(warehouse.FieldID, field.TypeUUID))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Warehouse.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, warehouse.FieldID)
		for _, f := range fields {
			if !warehouse.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != warehouse.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Name(); ok {
		_spec.SetField(warehouse.FieldName, field.TypeString, value)
	}
	if value, ok := wuo.mutation.LastUpdate(); ok {
		_spec.SetField(warehouse.FieldLastUpdate, field.TypeTime, value)
	}
	if wuo.mutation.LastUpdateCleared() {
		_spec.ClearField(warehouse.FieldLastUpdate, field.TypeTime)
	}
	if value, ok := wuo.mutation.OriginalData(); ok {
		_spec.SetField(warehouse.FieldOriginalData, field.TypeString, value)
	}
	if wuo.mutation.OriginalDataCleared() {
		_spec.ClearField(warehouse.FieldOriginalData, field.TypeString)
	}
	if value, ok := wuo.mutation.Enabled(); ok {
		_spec.SetField(warehouse.FieldEnabled, field.TypeBool, value)
	}
	if value, ok := wuo.mutation.Filters(); ok {
		_spec.SetField(warehouse.FieldFilters, field.TypeJSON, value)
	}
	if value, ok := wuo.mutation.AppendedFilters(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, warehouse.FieldFilters, value)
		})
	}
	if wuo.mutation.FiltersCleared() {
		_spec.ClearField(warehouse.FieldFilters, field.TypeJSON)
	}
	if wuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   warehouse.ProductsTable,
			Columns: []string{warehouse.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !wuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   warehouse.ProductsTable,
			Columns: []string{warehouse.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   warehouse.ProductsTable,
			Columns: []string{warehouse.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.VendorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   warehouse.VendorTable,
			Columns: []string{warehouse.VendorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vendor.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.VendorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   warehouse.VendorTable,
			Columns: []string{warehouse.VendorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vendor.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Warehouse{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{warehouse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}
