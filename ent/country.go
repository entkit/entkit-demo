// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/entkit/entkit-demo/ent/country"
	"github.com/google/uuid"
)

// Country is the model entity for the Country schema.
type Country struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CountryQuery when eager-loading is set.
	Edges        CountryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CountryEdges holds the relations/edges for other nodes in the graph.
type CountryEdges struct {
	// Companies holds the value of the companies edge.
	Companies []*Company `json:"companies,omitempty"`
	// Phones holds the value of the phones edge.
	Phones []*Phone `json:"phones,omitempty"`
	// Emails holds the value of the emails edge.
	Emails []*Email `json:"emails,omitempty"`
	// Websites holds the value of the websites edge.
	Websites []*Website `json:"websites,omitempty"`
	// Locations holds the value of the locations edge.
	Locations []*Location `json:"locations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
	// totalCount holds the count of the edges above.
	totalCount [5]map[string]int

	namedCompanies map[string][]*Company
	namedPhones    map[string][]*Phone
	namedEmails    map[string][]*Email
	namedWebsites  map[string][]*Website
	namedLocations map[string][]*Location
}

// CompaniesOrErr returns the Companies value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) CompaniesOrErr() ([]*Company, error) {
	if e.loadedTypes[0] {
		return e.Companies, nil
	}
	return nil, &NotLoadedError{edge: "companies"}
}

// PhonesOrErr returns the Phones value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) PhonesOrErr() ([]*Phone, error) {
	if e.loadedTypes[1] {
		return e.Phones, nil
	}
	return nil, &NotLoadedError{edge: "phones"}
}

// EmailsOrErr returns the Emails value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) EmailsOrErr() ([]*Email, error) {
	if e.loadedTypes[2] {
		return e.Emails, nil
	}
	return nil, &NotLoadedError{edge: "emails"}
}

// WebsitesOrErr returns the Websites value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) WebsitesOrErr() ([]*Website, error) {
	if e.loadedTypes[3] {
		return e.Websites, nil
	}
	return nil, &NotLoadedError{edge: "websites"}
}

// LocationsOrErr returns the Locations value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) LocationsOrErr() ([]*Location, error) {
	if e.loadedTypes[4] {
		return e.Locations, nil
	}
	return nil, &NotLoadedError{edge: "locations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Country) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case country.FieldName, country.FieldCode:
			values[i] = new(sql.NullString)
		case country.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Country fields.
func (c *Country) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case country.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case country.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case country.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				c.Code = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Country.
// This includes values selected through modifiers, order, etc.
func (c *Country) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryCompanies queries the "companies" edge of the Country entity.
func (c *Country) QueryCompanies() *CompanyQuery {
	return NewCountryClient(c.config).QueryCompanies(c)
}

// QueryPhones queries the "phones" edge of the Country entity.
func (c *Country) QueryPhones() *PhoneQuery {
	return NewCountryClient(c.config).QueryPhones(c)
}

// QueryEmails queries the "emails" edge of the Country entity.
func (c *Country) QueryEmails() *EmailQuery {
	return NewCountryClient(c.config).QueryEmails(c)
}

// QueryWebsites queries the "websites" edge of the Country entity.
func (c *Country) QueryWebsites() *WebsiteQuery {
	return NewCountryClient(c.config).QueryWebsites(c)
}

// QueryLocations queries the "locations" edge of the Country entity.
func (c *Country) QueryLocations() *LocationQuery {
	return NewCountryClient(c.config).QueryLocations(c)
}

// Update returns a builder for updating this Country.
// Note that you need to call Country.Unwrap() before calling this method if this Country
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Country) Update() *CountryUpdateOne {
	return NewCountryClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Country entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Country) Unwrap() *Country {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Country is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Country) String() string {
	var builder strings.Builder
	builder.WriteString("Country(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(c.Code)
	builder.WriteByte(')')
	return builder.String()
}

// NamedCompanies returns the Companies named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedCompanies(name string) ([]*Company, error) {
	if c.Edges.namedCompanies == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedCompanies[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedCompanies(name string, edges ...*Company) {
	if c.Edges.namedCompanies == nil {
		c.Edges.namedCompanies = make(map[string][]*Company)
	}
	if len(edges) == 0 {
		c.Edges.namedCompanies[name] = []*Company{}
	} else {
		c.Edges.namedCompanies[name] = append(c.Edges.namedCompanies[name], edges...)
	}
}

// NamedPhones returns the Phones named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedPhones(name string) ([]*Phone, error) {
	if c.Edges.namedPhones == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedPhones[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedPhones(name string, edges ...*Phone) {
	if c.Edges.namedPhones == nil {
		c.Edges.namedPhones = make(map[string][]*Phone)
	}
	if len(edges) == 0 {
		c.Edges.namedPhones[name] = []*Phone{}
	} else {
		c.Edges.namedPhones[name] = append(c.Edges.namedPhones[name], edges...)
	}
}

// NamedEmails returns the Emails named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedEmails(name string) ([]*Email, error) {
	if c.Edges.namedEmails == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedEmails[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedEmails(name string, edges ...*Email) {
	if c.Edges.namedEmails == nil {
		c.Edges.namedEmails = make(map[string][]*Email)
	}
	if len(edges) == 0 {
		c.Edges.namedEmails[name] = []*Email{}
	} else {
		c.Edges.namedEmails[name] = append(c.Edges.namedEmails[name], edges...)
	}
}

// NamedWebsites returns the Websites named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedWebsites(name string) ([]*Website, error) {
	if c.Edges.namedWebsites == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedWebsites[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedWebsites(name string, edges ...*Website) {
	if c.Edges.namedWebsites == nil {
		c.Edges.namedWebsites = make(map[string][]*Website)
	}
	if len(edges) == 0 {
		c.Edges.namedWebsites[name] = []*Website{}
	} else {
		c.Edges.namedWebsites[name] = append(c.Edges.namedWebsites[name], edges...)
	}
}

// NamedLocations returns the Locations named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedLocations(name string) ([]*Location, error) {
	if c.Edges.namedLocations == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedLocations[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedLocations(name string, edges ...*Location) {
	if c.Edges.namedLocations == nil {
		c.Edges.namedLocations = make(map[string][]*Location)
	}
	if len(edges) == 0 {
		c.Edges.namedLocations[name] = []*Location{}
	} else {
		c.Edges.namedLocations[name] = append(c.Edges.namedLocations[name], edges...)
	}
}

// Countries is a parsable slice of Country.
type Countries []*Country
