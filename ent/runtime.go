// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/entkit/entkit-demo/ent/company"
	"github.com/entkit/entkit-demo/ent/country"
	"github.com/entkit/entkit-demo/ent/email"
	"github.com/entkit/entkit-demo/ent/image"
	"github.com/entkit/entkit-demo/ent/location"
	"github.com/entkit/entkit-demo/ent/phone"
	"github.com/entkit/entkit-demo/ent/product"
	"github.com/entkit/entkit-demo/ent/schema"
	"github.com/entkit/entkit-demo/ent/vendor"
	"github.com/entkit/entkit-demo/ent/warehouse"
	"github.com/entkit/entkit-demo/ent/website"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	companyFields := schema.Company{}.Fields()
	_ = companyFields
	// companyDescName is the schema descriptor for name field.
	companyDescName := companyFields[1].Descriptor()
	// company.NameValidator is a validator for the "name" field. It is called by the builders before save.
	company.NameValidator = companyDescName.Validators[0].(func(string) error)
	// companyDescDescription is the schema descriptor for description field.
	companyDescDescription := companyFields[2].Descriptor()
	// company.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	company.DescriptionValidator = companyDescDescription.Validators[0].(func(string) error)
	// companyDescID is the schema descriptor for id field.
	companyDescID := companyFields[0].Descriptor()
	// company.DefaultID holds the default value on creation for the id field.
	company.DefaultID = companyDescID.Default.(func() uuid.UUID)
	countryFields := schema.Country{}.Fields()
	_ = countryFields
	// countryDescName is the schema descriptor for name field.
	countryDescName := countryFields[1].Descriptor()
	// country.NameValidator is a validator for the "name" field. It is called by the builders before save.
	country.NameValidator = countryDescName.Validators[0].(func(string) error)
	// countryDescCode is the schema descriptor for code field.
	countryDescCode := countryFields[2].Descriptor()
	// country.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	country.CodeValidator = countryDescCode.Validators[0].(func(string) error)
	// countryDescID is the schema descriptor for id field.
	countryDescID := countryFields[0].Descriptor()
	// country.DefaultID holds the default value on creation for the id field.
	country.DefaultID = countryDescID.Default.(func() uuid.UUID)
	emailFields := schema.Email{}.Fields()
	_ = emailFields
	// emailDescTitle is the schema descriptor for title field.
	emailDescTitle := emailFields[1].Descriptor()
	// email.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	email.TitleValidator = emailDescTitle.Validators[0].(func(string) error)
	// emailDescDescription is the schema descriptor for description field.
	emailDescDescription := emailFields[2].Descriptor()
	// email.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	email.DescriptionValidator = emailDescDescription.Validators[0].(func(string) error)
	// emailDescAddress is the schema descriptor for address field.
	emailDescAddress := emailFields[3].Descriptor()
	// email.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	email.AddressValidator = emailDescAddress.Validators[0].(func(string) error)
	// emailDescID is the schema descriptor for id field.
	emailDescID := emailFields[0].Descriptor()
	// email.DefaultID holds the default value on creation for the id field.
	email.DefaultID = emailDescID.Default.(func() uuid.UUID)
	imageFields := schema.Image{}.Fields()
	_ = imageFields
	// imageDescTitle is the schema descriptor for title field.
	imageDescTitle := imageFields[1].Descriptor()
	// image.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	image.TitleValidator = imageDescTitle.Validators[0].(func(string) error)
	// imageDescOriginalURL is the schema descriptor for original_url field.
	imageDescOriginalURL := imageFields[2].Descriptor()
	// image.OriginalURLValidator is a validator for the "original_url" field. It is called by the builders before save.
	image.OriginalURLValidator = imageDescOriginalURL.Validators[0].(func(string) error)
	// imageDescID is the schema descriptor for id field.
	imageDescID := imageFields[0].Descriptor()
	// image.DefaultID holds the default value on creation for the id field.
	image.DefaultID = imageDescID.Default.(func() uuid.UUID)
	locationFields := schema.Location{}.Fields()
	_ = locationFields
	// locationDescTitle is the schema descriptor for title field.
	locationDescTitle := locationFields[1].Descriptor()
	// location.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	location.TitleValidator = locationDescTitle.Validators[0].(func(string) error)
	// locationDescDescription is the schema descriptor for description field.
	locationDescDescription := locationFields[2].Descriptor()
	// location.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	location.DescriptionValidator = locationDescDescription.Validators[0].(func(string) error)
	// locationDescAddress is the schema descriptor for address field.
	locationDescAddress := locationFields[5].Descriptor()
	// location.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	location.AddressValidator = locationDescAddress.Validators[0].(func(string) error)
	// locationDescPostcode is the schema descriptor for postcode field.
	locationDescPostcode := locationFields[6].Descriptor()
	// location.PostcodeValidator is a validator for the "postcode" field. It is called by the builders before save.
	location.PostcodeValidator = locationDescPostcode.Validators[0].(func(string) error)
	// locationDescType is the schema descriptor for type field.
	locationDescType := locationFields[7].Descriptor()
	// location.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	location.TypeValidator = locationDescType.Validators[0].(func(string) error)
	// locationDescState is the schema descriptor for state field.
	locationDescState := locationFields[8].Descriptor()
	// location.StateValidator is a validator for the "state" field. It is called by the builders before save.
	location.StateValidator = locationDescState.Validators[0].(func(string) error)
	// locationDescSuburb is the schema descriptor for suburb field.
	locationDescSuburb := locationFields[9].Descriptor()
	// location.SuburbValidator is a validator for the "suburb" field. It is called by the builders before save.
	location.SuburbValidator = locationDescSuburb.Validators[0].(func(string) error)
	// locationDescStreetType is the schema descriptor for street_type field.
	locationDescStreetType := locationFields[10].Descriptor()
	// location.StreetTypeValidator is a validator for the "street_type" field. It is called by the builders before save.
	location.StreetTypeValidator = locationDescStreetType.Validators[0].(func(string) error)
	// locationDescStreetName is the schema descriptor for street_name field.
	locationDescStreetName := locationFields[11].Descriptor()
	// location.StreetNameValidator is a validator for the "street_name" field. It is called by the builders before save.
	location.StreetNameValidator = locationDescStreetName.Validators[0].(func(string) error)
	// locationDescID is the schema descriptor for id field.
	locationDescID := locationFields[0].Descriptor()
	// location.DefaultID holds the default value on creation for the id field.
	location.DefaultID = locationDescID.Default.(func() uuid.UUID)
	phoneFields := schema.Phone{}.Fields()
	_ = phoneFields
	// phoneDescTitle is the schema descriptor for title field.
	phoneDescTitle := phoneFields[1].Descriptor()
	// phone.DefaultTitle holds the default value on creation for the title field.
	phone.DefaultTitle = phoneDescTitle.Default.(string)
	// phone.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	phone.TitleValidator = phoneDescTitle.Validators[0].(func(string) error)
	// phoneDescDescription is the schema descriptor for description field.
	phoneDescDescription := phoneFields[2].Descriptor()
	// phone.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	phone.DescriptionValidator = phoneDescDescription.Validators[0].(func(string) error)
	// phoneDescNumber is the schema descriptor for number field.
	phoneDescNumber := phoneFields[3].Descriptor()
	// phone.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	phone.NumberValidator = phoneDescNumber.Validators[0].(func(string) error)
	// phoneDescType is the schema descriptor for type field.
	phoneDescType := phoneFields[4].Descriptor()
	// phone.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	phone.TypeValidator = phoneDescType.Validators[0].(func(string) error)
	// phoneDescID is the schema descriptor for id field.
	phoneDescID := phoneFields[0].Descriptor()
	// phone.DefaultID holds the default value on creation for the id field.
	phone.DefaultID = phoneDescID.Default.(func() uuid.UUID)
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescName is the schema descriptor for name field.
	productDescName := productFields[1].Descriptor()
	// product.NameValidator is a validator for the "name" field. It is called by the builders before save.
	product.NameValidator = productDescName.Validators[0].(func(string) error)
	// productDescDescription is the schema descriptor for description field.
	productDescDescription := productFields[2].Descriptor()
	// product.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	product.DescriptionValidator = productDescDescription.Validators[0].(func(string) error)
	// productDescImage is the schema descriptor for image field.
	productDescImage := productFields[3].Descriptor()
	// product.ImageValidator is a validator for the "image" field. It is called by the builders before save.
	product.ImageValidator = productDescImage.Validators[0].(func(string) error)
	// productDescURL is the schema descriptor for url field.
	productDescURL := productFields[4].Descriptor()
	// product.URLValidator is a validator for the "url" field. It is called by the builders before save.
	product.URLValidator = productDescURL.Validators[0].(func(string) error)
	// productDescID is the schema descriptor for id field.
	productDescID := productFields[0].Descriptor()
	// product.DefaultID holds the default value on creation for the id field.
	product.DefaultID = productDescID.Default.(func() uuid.UUID)
	vendorFields := schema.Vendor{}.Fields()
	_ = vendorFields
	// vendorDescName is the schema descriptor for name field.
	vendorDescName := vendorFields[1].Descriptor()
	// vendor.NameValidator is a validator for the "name" field. It is called by the builders before save.
	vendor.NameValidator = vendorDescName.Validators[0].(func(string) error)
	// vendorDescID is the schema descriptor for id field.
	vendorDescID := vendorFields[0].Descriptor()
	// vendor.DefaultID holds the default value on creation for the id field.
	vendor.DefaultID = vendorDescID.Default.(func() uuid.UUID)
	warehouseFields := schema.Warehouse{}.Fields()
	_ = warehouseFields
	// warehouseDescName is the schema descriptor for name field.
	warehouseDescName := warehouseFields[1].Descriptor()
	// warehouse.NameValidator is a validator for the "name" field. It is called by the builders before save.
	warehouse.NameValidator = warehouseDescName.Validators[0].(func(string) error)
	// warehouseDescEnabled is the schema descriptor for enabled field.
	warehouseDescEnabled := warehouseFields[4].Descriptor()
	// warehouse.DefaultEnabled holds the default value on creation for the enabled field.
	warehouse.DefaultEnabled = warehouseDescEnabled.Default.(bool)
	// warehouseDescID is the schema descriptor for id field.
	warehouseDescID := warehouseFields[0].Descriptor()
	// warehouse.DefaultID holds the default value on creation for the id field.
	warehouse.DefaultID = warehouseDescID.Default.(func() uuid.UUID)
	websiteFields := schema.Website{}.Fields()
	_ = websiteFields
	// websiteDescTitle is the schema descriptor for title field.
	websiteDescTitle := websiteFields[1].Descriptor()
	// website.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	website.TitleValidator = websiteDescTitle.Validators[0].(func(string) error)
	// websiteDescDescription is the schema descriptor for description field.
	websiteDescDescription := websiteFields[2].Descriptor()
	// website.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	website.DescriptionValidator = websiteDescDescription.Validators[0].(func(string) error)
	// websiteDescURL is the schema descriptor for url field.
	websiteDescURL := websiteFields[3].Descriptor()
	// website.URLValidator is a validator for the "url" field. It is called by the builders before save.
	website.URLValidator = websiteDescURL.Validators[0].(func(string) error)
	// websiteDescID is the schema descriptor for id field.
	websiteDescID := websiteFields[0].Descriptor()
	// website.DefaultID holds the default value on creation for the id field.
	website.DefaultID = websiteDescID.Default.(func() uuid.UUID)
}
