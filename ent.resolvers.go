package entproject

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/entkit/entkit"
	"github.com/entkit/entkit-demo/ent"
	"github.com/google/uuid"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

// Companies is the resolver for the companies field.
func (r *queryResolver) Companies(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy *ent.CompanyOrder, where *ent.CompanyWhereInput, q *string) (*ent.CompanyConnection, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Company", entkit.ActionReadScope)
	if err != nil {
		return nil, err
	}
	return r.client.Company.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithCompanyOrder(orderBy),
			ent.WithCompanyFilter(where.ApplySearchQuery(q).Filter),
		)
}

// Countries is the resolver for the countries field.
func (r *queryResolver) Countries(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy *ent.CountryOrder, where *ent.CountryWhereInput, q *string) (*ent.CountryConnection, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Country", entkit.ActionReadScope)
	if err != nil {
		return nil, err
	}
	return r.client.Country.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithCountryOrder(orderBy),
			ent.WithCountryFilter(where.ApplySearchQuery(q).Filter),
		)
}

// Emails is the resolver for the emails field.
func (r *queryResolver) Emails(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy *ent.EmailOrder, where *ent.EmailWhereInput, q *string) (*ent.EmailConnection, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Email", entkit.ActionReadScope)
	if err != nil {
		return nil, err
	}
	return r.client.Email.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithEmailOrder(orderBy),
			ent.WithEmailFilter(where.ApplySearchQuery(q).Filter),
		)
}

// Images is the resolver for the images field.
func (r *queryResolver) Images(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy *ent.ImageOrder, where *ent.ImageWhereInput, q *string) (*ent.ImageConnection, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Image", entkit.ActionReadScope)
	if err != nil {
		return nil, err
	}
	return r.client.Image.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithImageOrder(orderBy),
			ent.WithImageFilter(where.ApplySearchQuery(q).Filter),
		)
}

// Locations is the resolver for the locations field.
func (r *queryResolver) Locations(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy *ent.LocationOrder, where *ent.LocationWhereInput, q *string) (*ent.LocationConnection, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Location", entkit.ActionReadScope)
	if err != nil {
		return nil, err
	}
	return r.client.Location.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithLocationOrder(orderBy),
			ent.WithLocationFilter(where.ApplySearchQuery(q).Filter),
		)
}

// Phones is the resolver for the phones field.
func (r *queryResolver) Phones(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy *ent.PhoneOrder, where *ent.PhoneWhereInput, q *string) (*ent.PhoneConnection, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Phone", entkit.ActionReadScope)
	if err != nil {
		return nil, err
	}
	return r.client.Phone.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithPhoneOrder(orderBy),
			ent.WithPhoneFilter(where.ApplySearchQuery(q).Filter),
		)
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy *ent.ProductOrder, where *ent.ProductWhereInput, q *string) (*ent.ProductConnection, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Product", entkit.ActionReadScope)
	if err != nil {
		return nil, err
	}
	return r.client.Product.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithProductOrder(orderBy),
			ent.WithProductFilter(where.ApplySearchQuery(q).Filter),
		)
}

// Vendors is the resolver for the vendors field.
func (r *queryResolver) Vendors(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy *ent.VendorOrder, where *ent.VendorWhereInput, q *string) (*ent.VendorConnection, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Vendor", entkit.ActionReadScope)
	if err != nil {
		return nil, err
	}
	return r.client.Vendor.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithVendorOrder(orderBy),
			ent.WithVendorFilter(where.ApplySearchQuery(q).Filter),
		)
}

// Warehouses is the resolver for the warehouses field.
func (r *queryResolver) Warehouses(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy *ent.WarehouseOrder, where *ent.WarehouseWhereInput, q *string) (*ent.WarehouseConnection, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Warehouse", entkit.ActionReadScope)
	if err != nil {
		return nil, err
	}
	return r.client.Warehouse.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithWarehouseOrder(orderBy),
			ent.WithWarehouseFilter(where.ApplySearchQuery(q).Filter),
		)
}

// Websites is the resolver for the websites field.
func (r *queryResolver) Websites(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy *ent.WebsiteOrder, where *ent.WebsiteWhereInput, q *string) (*ent.WebsiteConnection, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Website", entkit.ActionReadScope)
	if err != nil {
		return nil, err
	}
	return r.client.Website.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithWebsiteOrder(orderBy),
			ent.WithWebsiteFilter(where.ApplySearchQuery(q).Filter),
		)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
