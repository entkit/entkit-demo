// Autogenerated by entkit
//
// Code generated by entc, DO NOT EDIT.

package country

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/entkit/entkit-demo/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldName, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCode, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldName, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Country {
	return predicate.Country(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Country {
	return predicate.Country(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Country {
	return predicate.Country(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Country {
	return predicate.Country(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Country {
	return predicate.Country(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Country {
	return predicate.Country(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Country {
	return predicate.Country(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Country {
	return predicate.Country(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Country {
	return predicate.Country(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Country {
	return predicate.Country(sql.FieldContainsFold(FieldCode, v))
}

// HasCompanies applies the HasEdge predicate on the "companies" edge.
func HasCompanies() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CompaniesTable, CompaniesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompaniesWith applies the HasEdge predicate on the "companies" edge with a given conditions (other predicates).
func HasCompaniesWith(preds ...predicate.Company) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := newCompaniesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPhones applies the HasEdge predicate on the "phones" edge.
func HasPhones() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PhonesTable, PhonesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPhonesWith applies the HasEdge predicate on the "phones" edge with a given conditions (other predicates).
func HasPhonesWith(preds ...predicate.Phone) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := newPhonesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEmails applies the HasEdge predicate on the "emails" edge.
func HasEmails() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EmailsTable, EmailsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmailsWith applies the HasEdge predicate on the "emails" edge with a given conditions (other predicates).
func HasEmailsWith(preds ...predicate.Email) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := newEmailsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWebsites applies the HasEdge predicate on the "websites" edge.
func HasWebsites() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WebsitesTable, WebsitesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWebsitesWith applies the HasEdge predicate on the "websites" edge with a given conditions (other predicates).
func HasWebsitesWith(preds ...predicate.Website) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := newWebsitesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLocations applies the HasEdge predicate on the "locations" edge.
func HasLocations() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LocationsTable, LocationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocationsWith applies the HasEdge predicate on the "locations" edge with a given conditions (other predicates).
func HasLocationsWith(preds ...predicate.Location) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := newLocationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Country) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Country) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Country) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		p(s.Not())
	})
}