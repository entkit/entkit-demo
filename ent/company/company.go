// Code generated by ent, DO NOT EDIT.

package company

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the company type in the database.
	Label = "company"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeCountries holds the string denoting the countries edge name in mutations.
	EdgeCountries = "countries"
	// EdgePhones holds the string denoting the phones edge name in mutations.
	EdgePhones = "phones"
	// EdgeEmails holds the string denoting the emails edge name in mutations.
	EdgeEmails = "emails"
	// EdgeWebsites holds the string denoting the websites edge name in mutations.
	EdgeWebsites = "websites"
	// EdgeLocations holds the string denoting the locations edge name in mutations.
	EdgeLocations = "locations"
	// EdgeLogoImage holds the string denoting the logo_image edge name in mutations.
	EdgeLogoImage = "logo_image"
	// EdgeCoverImage holds the string denoting the cover_image edge name in mutations.
	EdgeCoverImage = "cover_image"
	// EdgeGalleryImages holds the string denoting the gallery_images edge name in mutations.
	EdgeGalleryImages = "gallery_images"
	// Table holds the table name of the company in the database.
	Table = "companies"
	// CountriesTable is the table that holds the countries relation/edge. The primary key declared below.
	CountriesTable = "company_countries"
	// CountriesInverseTable is the table name for the Country entity.
	// It exists in this package in order to avoid circular dependency with the "country" package.
	CountriesInverseTable = "countries"
	// PhonesTable is the table that holds the phones relation/edge.
	PhonesTable = "phones"
	// PhonesInverseTable is the table name for the Phone entity.
	// It exists in this package in order to avoid circular dependency with the "phone" package.
	PhonesInverseTable = "phones"
	// PhonesColumn is the table column denoting the phones relation/edge.
	PhonesColumn = "company_phones"
	// EmailsTable is the table that holds the emails relation/edge.
	EmailsTable = "emails"
	// EmailsInverseTable is the table name for the Email entity.
	// It exists in this package in order to avoid circular dependency with the "email" package.
	EmailsInverseTable = "emails"
	// EmailsColumn is the table column denoting the emails relation/edge.
	EmailsColumn = "company_emails"
	// WebsitesTable is the table that holds the websites relation/edge.
	WebsitesTable = "websites"
	// WebsitesInverseTable is the table name for the Website entity.
	// It exists in this package in order to avoid circular dependency with the "website" package.
	WebsitesInverseTable = "websites"
	// WebsitesColumn is the table column denoting the websites relation/edge.
	WebsitesColumn = "company_websites"
	// LocationsTable is the table that holds the locations relation/edge.
	LocationsTable = "locations"
	// LocationsInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationsInverseTable = "locations"
	// LocationsColumn is the table column denoting the locations relation/edge.
	LocationsColumn = "company_locations"
	// LogoImageTable is the table that holds the logo_image relation/edge.
	LogoImageTable = "images"
	// LogoImageInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	LogoImageInverseTable = "images"
	// LogoImageColumn is the table column denoting the logo_image relation/edge.
	LogoImageColumn = "company_logo_image"
	// CoverImageTable is the table that holds the cover_image relation/edge.
	CoverImageTable = "images"
	// CoverImageInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	CoverImageInverseTable = "images"
	// CoverImageColumn is the table column denoting the cover_image relation/edge.
	CoverImageColumn = "company_cover_image"
	// GalleryImagesTable is the table that holds the gallery_images relation/edge.
	GalleryImagesTable = "images"
	// GalleryImagesInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	GalleryImagesInverseTable = "images"
	// GalleryImagesColumn is the table column denoting the gallery_images relation/edge.
	GalleryImagesColumn = "company_gallery_images"
)

// Columns holds all SQL columns for company fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
}

var (
	// CountriesPrimaryKey and CountriesColumn2 are the table columns denoting the
	// primary key for the countries relation (M2M).
	CountriesPrimaryKey = []string{"company_id", "country_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Company queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCountriesCount orders the results by countries count.
func ByCountriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCountriesStep(), opts...)
	}
}

// ByCountries orders the results by countries terms.
func ByCountries(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCountriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPhonesCount orders the results by phones count.
func ByPhonesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPhonesStep(), opts...)
	}
}

// ByPhones orders the results by phones terms.
func ByPhones(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPhonesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEmailsCount orders the results by emails count.
func ByEmailsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEmailsStep(), opts...)
	}
}

// ByEmails orders the results by emails terms.
func ByEmails(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEmailsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWebsitesCount orders the results by websites count.
func ByWebsitesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWebsitesStep(), opts...)
	}
}

// ByWebsites orders the results by websites terms.
func ByWebsites(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWebsitesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLocationsCount orders the results by locations count.
func ByLocationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLocationsStep(), opts...)
	}
}

// ByLocations orders the results by locations terms.
func ByLocations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLocationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLogoImageField orders the results by logo_image field.
func ByLogoImageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLogoImageStep(), sql.OrderByField(field, opts...))
	}
}

// ByCoverImageField orders the results by cover_image field.
func ByCoverImageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCoverImageStep(), sql.OrderByField(field, opts...))
	}
}

// ByGalleryImagesCount orders the results by gallery_images count.
func ByGalleryImagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGalleryImagesStep(), opts...)
	}
}

// ByGalleryImages orders the results by gallery_images terms.
func ByGalleryImages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGalleryImagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCountriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CountriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, CountriesTable, CountriesPrimaryKey...),
	)
}
func newPhonesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PhonesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PhonesTable, PhonesColumn),
	)
}
func newEmailsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EmailsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EmailsTable, EmailsColumn),
	)
}
func newWebsitesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WebsitesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WebsitesTable, WebsitesColumn),
	)
}
func newLocationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LocationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LocationsTable, LocationsColumn),
	)
}
func newLogoImageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LogoImageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, LogoImageTable, LogoImageColumn),
	)
}
func newCoverImageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CoverImageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, CoverImageTable, CoverImageColumn),
	)
}
func newGalleryImagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GalleryImagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, GalleryImagesTable, GalleryImagesColumn),
	)
}
