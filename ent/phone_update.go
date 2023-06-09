// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/entkit/entkit-demo/ent/company"
	"github.com/entkit/entkit-demo/ent/country"
	"github.com/entkit/entkit-demo/ent/phone"
	"github.com/entkit/entkit-demo/ent/predicate"
	"github.com/google/uuid"
)

// PhoneUpdate is the builder for updating Phone entities.
type PhoneUpdate struct {
	config
	hooks    []Hook
	mutation *PhoneMutation
}

// Where appends a list predicates to the PhoneUpdate builder.
func (pu *PhoneUpdate) Where(ps ...predicate.Phone) *PhoneUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetTitle sets the "title" field.
func (pu *PhoneUpdate) SetTitle(s string) *PhoneUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pu *PhoneUpdate) SetNillableTitle(s *string) *PhoneUpdate {
	if s != nil {
		pu.SetTitle(*s)
	}
	return pu
}

// SetDescription sets the "description" field.
func (pu *PhoneUpdate) SetDescription(s string) *PhoneUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNumber sets the "number" field.
func (pu *PhoneUpdate) SetNumber(s string) *PhoneUpdate {
	pu.mutation.SetNumber(s)
	return pu
}

// SetType sets the "type" field.
func (pu *PhoneUpdate) SetType(s string) *PhoneUpdate {
	pu.mutation.SetType(s)
	return pu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (pu *PhoneUpdate) SetCompanyID(id uuid.UUID) *PhoneUpdate {
	pu.mutation.SetCompanyID(id)
	return pu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (pu *PhoneUpdate) SetNillableCompanyID(id *uuid.UUID) *PhoneUpdate {
	if id != nil {
		pu = pu.SetCompanyID(*id)
	}
	return pu
}

// SetCompany sets the "company" edge to the Company entity.
func (pu *PhoneUpdate) SetCompany(c *Company) *PhoneUpdate {
	return pu.SetCompanyID(c.ID)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (pu *PhoneUpdate) SetCountryID(id uuid.UUID) *PhoneUpdate {
	pu.mutation.SetCountryID(id)
	return pu
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (pu *PhoneUpdate) SetNillableCountryID(id *uuid.UUID) *PhoneUpdate {
	if id != nil {
		pu = pu.SetCountryID(*id)
	}
	return pu
}

// SetCountry sets the "country" edge to the Country entity.
func (pu *PhoneUpdate) SetCountry(c *Country) *PhoneUpdate {
	return pu.SetCountryID(c.ID)
}

// Mutation returns the PhoneMutation object of the builder.
func (pu *PhoneUpdate) Mutation() *PhoneMutation {
	return pu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (pu *PhoneUpdate) ClearCompany() *PhoneUpdate {
	pu.mutation.ClearCompany()
	return pu
}

// ClearCountry clears the "country" edge to the Country entity.
func (pu *PhoneUpdate) ClearCountry() *PhoneUpdate {
	pu.mutation.ClearCountry()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PhoneUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PhoneMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PhoneUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PhoneUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PhoneUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PhoneUpdate) check() error {
	if v, ok := pu.mutation.Title(); ok {
		if err := phone.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Phone.title": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Description(); ok {
		if err := phone.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Phone.description": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Number(); ok {
		if err := phone.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Phone.number": %w`, err)}
		}
	}
	if v, ok := pu.mutation.GetType(); ok {
		if err := phone.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Phone.type": %w`, err)}
		}
	}
	return nil
}

func (pu *PhoneUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(phone.Table, phone.Columns, sqlgraph.NewFieldSpec(phone.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(phone.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(phone.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.Number(); ok {
		_spec.SetField(phone.FieldNumber, field.TypeString, value)
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.SetField(phone.FieldType, field.TypeString, value)
	}
	if pu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phone.CompanyTable,
			Columns: []string{phone.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phone.CompanyTable,
			Columns: []string{phone.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phone.CountryTable,
			Columns: []string{phone.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phone.CountryTable,
			Columns: []string{phone.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{phone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PhoneUpdateOne is the builder for updating a single Phone entity.
type PhoneUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PhoneMutation
}

// SetTitle sets the "title" field.
func (puo *PhoneUpdateOne) SetTitle(s string) *PhoneUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (puo *PhoneUpdateOne) SetNillableTitle(s *string) *PhoneUpdateOne {
	if s != nil {
		puo.SetTitle(*s)
	}
	return puo
}

// SetDescription sets the "description" field.
func (puo *PhoneUpdateOne) SetDescription(s string) *PhoneUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNumber sets the "number" field.
func (puo *PhoneUpdateOne) SetNumber(s string) *PhoneUpdateOne {
	puo.mutation.SetNumber(s)
	return puo
}

// SetType sets the "type" field.
func (puo *PhoneUpdateOne) SetType(s string) *PhoneUpdateOne {
	puo.mutation.SetType(s)
	return puo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (puo *PhoneUpdateOne) SetCompanyID(id uuid.UUID) *PhoneUpdateOne {
	puo.mutation.SetCompanyID(id)
	return puo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (puo *PhoneUpdateOne) SetNillableCompanyID(id *uuid.UUID) *PhoneUpdateOne {
	if id != nil {
		puo = puo.SetCompanyID(*id)
	}
	return puo
}

// SetCompany sets the "company" edge to the Company entity.
func (puo *PhoneUpdateOne) SetCompany(c *Company) *PhoneUpdateOne {
	return puo.SetCompanyID(c.ID)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (puo *PhoneUpdateOne) SetCountryID(id uuid.UUID) *PhoneUpdateOne {
	puo.mutation.SetCountryID(id)
	return puo
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (puo *PhoneUpdateOne) SetNillableCountryID(id *uuid.UUID) *PhoneUpdateOne {
	if id != nil {
		puo = puo.SetCountryID(*id)
	}
	return puo
}

// SetCountry sets the "country" edge to the Country entity.
func (puo *PhoneUpdateOne) SetCountry(c *Country) *PhoneUpdateOne {
	return puo.SetCountryID(c.ID)
}

// Mutation returns the PhoneMutation object of the builder.
func (puo *PhoneUpdateOne) Mutation() *PhoneMutation {
	return puo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (puo *PhoneUpdateOne) ClearCompany() *PhoneUpdateOne {
	puo.mutation.ClearCompany()
	return puo
}

// ClearCountry clears the "country" edge to the Country entity.
func (puo *PhoneUpdateOne) ClearCountry() *PhoneUpdateOne {
	puo.mutation.ClearCountry()
	return puo
}

// Where appends a list predicates to the PhoneUpdate builder.
func (puo *PhoneUpdateOne) Where(ps ...predicate.Phone) *PhoneUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PhoneUpdateOne) Select(field string, fields ...string) *PhoneUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Phone entity.
func (puo *PhoneUpdateOne) Save(ctx context.Context) (*Phone, error) {
	return withHooks[*Phone, PhoneMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PhoneUpdateOne) SaveX(ctx context.Context) *Phone {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PhoneUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PhoneUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PhoneUpdateOne) check() error {
	if v, ok := puo.mutation.Title(); ok {
		if err := phone.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Phone.title": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Description(); ok {
		if err := phone.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Phone.description": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Number(); ok {
		if err := phone.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Phone.number": %w`, err)}
		}
	}
	if v, ok := puo.mutation.GetType(); ok {
		if err := phone.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Phone.type": %w`, err)}
		}
	}
	return nil
}

func (puo *PhoneUpdateOne) sqlSave(ctx context.Context) (_node *Phone, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(phone.Table, phone.Columns, sqlgraph.NewFieldSpec(phone.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Phone.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, phone.FieldID)
		for _, f := range fields {
			if !phone.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != phone.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(phone.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(phone.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.Number(); ok {
		_spec.SetField(phone.FieldNumber, field.TypeString, value)
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.SetField(phone.FieldType, field.TypeString, value)
	}
	if puo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phone.CompanyTable,
			Columns: []string{phone.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phone.CompanyTable,
			Columns: []string{phone.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phone.CountryTable,
			Columns: []string{phone.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phone.CountryTable,
			Columns: []string{phone.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Phone{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{phone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
