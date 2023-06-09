// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (c *Company) Countries(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *CountryOrder, where *CountryWhereInput,
) (*CountryConnection, error) {
	opts := []CountryPaginateOption{
		WithCountryOrder(orderBy),
		WithCountryFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[0][alias]
	if nodes, err := c.NamedCountries(alias); err == nil || hasTotalCount {
		pager, err := newCountryPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &CountryConnection{Edges: []*CountryEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryCountries().Paginate(ctx, after, first, before, last, opts...)
}

func (c *Company) Phones(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *PhoneOrder, where *PhoneWhereInput,
) (*PhoneConnection, error) {
	opts := []PhonePaginateOption{
		WithPhoneOrder(orderBy),
		WithPhoneFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[1][alias]
	if nodes, err := c.NamedPhones(alias); err == nil || hasTotalCount {
		pager, err := newPhonePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &PhoneConnection{Edges: []*PhoneEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryPhones().Paginate(ctx, after, first, before, last, opts...)
}

func (c *Company) Emails(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *EmailOrder, where *EmailWhereInput,
) (*EmailConnection, error) {
	opts := []EmailPaginateOption{
		WithEmailOrder(orderBy),
		WithEmailFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[2][alias]
	if nodes, err := c.NamedEmails(alias); err == nil || hasTotalCount {
		pager, err := newEmailPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &EmailConnection{Edges: []*EmailEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryEmails().Paginate(ctx, after, first, before, last, opts...)
}

func (c *Company) Websites(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *WebsiteOrder, where *WebsiteWhereInput,
) (*WebsiteConnection, error) {
	opts := []WebsitePaginateOption{
		WithWebsiteOrder(orderBy),
		WithWebsiteFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[3][alias]
	if nodes, err := c.NamedWebsites(alias); err == nil || hasTotalCount {
		pager, err := newWebsitePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &WebsiteConnection{Edges: []*WebsiteEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryWebsites().Paginate(ctx, after, first, before, last, opts...)
}

func (c *Company) Locations(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *LocationOrder, where *LocationWhereInput,
) (*LocationConnection, error) {
	opts := []LocationPaginateOption{
		WithLocationOrder(orderBy),
		WithLocationFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[4][alias]
	if nodes, err := c.NamedLocations(alias); err == nil || hasTotalCount {
		pager, err := newLocationPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &LocationConnection{Edges: []*LocationEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryLocations().Paginate(ctx, after, first, before, last, opts...)
}

func (c *Company) LogoImage(ctx context.Context) (*Image, error) {
	result, err := c.Edges.LogoImageOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryLogoImage().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Company) CoverImage(ctx context.Context) (*Image, error) {
	result, err := c.Edges.CoverImageOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryCoverImage().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Company) GalleryImages(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ImageOrder, where *ImageWhereInput,
) (*ImageConnection, error) {
	opts := []ImagePaginateOption{
		WithImageOrder(orderBy),
		WithImageFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[7][alias]
	if nodes, err := c.NamedGalleryImages(alias); err == nil || hasTotalCount {
		pager, err := newImagePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ImageConnection{Edges: []*ImageEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryGalleryImages().Paginate(ctx, after, first, before, last, opts...)
}

func (c *Country) Companies(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *CompanyOrder, where *CompanyWhereInput,
) (*CompanyConnection, error) {
	opts := []CompanyPaginateOption{
		WithCompanyOrder(orderBy),
		WithCompanyFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[0][alias]
	if nodes, err := c.NamedCompanies(alias); err == nil || hasTotalCount {
		pager, err := newCompanyPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &CompanyConnection{Edges: []*CompanyEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryCompanies().Paginate(ctx, after, first, before, last, opts...)
}

func (c *Country) Phones(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *PhoneOrder, where *PhoneWhereInput,
) (*PhoneConnection, error) {
	opts := []PhonePaginateOption{
		WithPhoneOrder(orderBy),
		WithPhoneFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[1][alias]
	if nodes, err := c.NamedPhones(alias); err == nil || hasTotalCount {
		pager, err := newPhonePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &PhoneConnection{Edges: []*PhoneEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryPhones().Paginate(ctx, after, first, before, last, opts...)
}

func (c *Country) Emails(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *EmailOrder, where *EmailWhereInput,
) (*EmailConnection, error) {
	opts := []EmailPaginateOption{
		WithEmailOrder(orderBy),
		WithEmailFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[2][alias]
	if nodes, err := c.NamedEmails(alias); err == nil || hasTotalCount {
		pager, err := newEmailPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &EmailConnection{Edges: []*EmailEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryEmails().Paginate(ctx, after, first, before, last, opts...)
}

func (c *Country) Websites(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *WebsiteOrder, where *WebsiteWhereInput,
) (*WebsiteConnection, error) {
	opts := []WebsitePaginateOption{
		WithWebsiteOrder(orderBy),
		WithWebsiteFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[3][alias]
	if nodes, err := c.NamedWebsites(alias); err == nil || hasTotalCount {
		pager, err := newWebsitePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &WebsiteConnection{Edges: []*WebsiteEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryWebsites().Paginate(ctx, after, first, before, last, opts...)
}

func (c *Country) Locations(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *LocationOrder, where *LocationWhereInput,
) (*LocationConnection, error) {
	opts := []LocationPaginateOption{
		WithLocationOrder(orderBy),
		WithLocationFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[4][alias]
	if nodes, err := c.NamedLocations(alias); err == nil || hasTotalCount {
		pager, err := newLocationPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &LocationConnection{Edges: []*LocationEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryLocations().Paginate(ctx, after, first, before, last, opts...)
}

func (e *Email) Company(ctx context.Context) (*Company, error) {
	result, err := e.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (e *Email) Country(ctx context.Context) (*Country, error) {
	result, err := e.Edges.CountryOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryCountry().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (i *Image) GalleryCompany(ctx context.Context) (*Company, error) {
	result, err := i.Edges.GalleryCompanyOrErr()
	if IsNotLoaded(err) {
		result, err = i.QueryGalleryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (i *Image) LogoCompany(ctx context.Context) (*Company, error) {
	result, err := i.Edges.LogoCompanyOrErr()
	if IsNotLoaded(err) {
		result, err = i.QueryLogoCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (i *Image) CoverCompany(ctx context.Context) (*Company, error) {
	result, err := i.Edges.CoverCompanyOrErr()
	if IsNotLoaded(err) {
		result, err = i.QueryCoverCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (l *Location) Company(ctx context.Context) (*Company, error) {
	result, err := l.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = l.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (l *Location) Country(ctx context.Context) (*Country, error) {
	result, err := l.Edges.CountryOrErr()
	if IsNotLoaded(err) {
		result, err = l.QueryCountry().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ph *Phone) Company(ctx context.Context) (*Company, error) {
	result, err := ph.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = ph.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ph *Phone) Country(ctx context.Context) (*Country, error) {
	result, err := ph.Edges.CountryOrErr()
	if IsNotLoaded(err) {
		result, err = ph.QueryCountry().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pr *Product) Warehouse(ctx context.Context) (*Warehouse, error) {
	result, err := pr.Edges.WarehouseOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryWarehouse().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pr *Product) Vendor(ctx context.Context) (*Vendor, error) {
	result, err := pr.Edges.VendorOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryVendor().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (v *Vendor) Warehouses(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *WarehouseOrder, where *WarehouseWhereInput,
) (*WarehouseConnection, error) {
	opts := []WarehousePaginateOption{
		WithWarehouseOrder(orderBy),
		WithWarehouseFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := v.Edges.totalCount[0][alias]
	if nodes, err := v.NamedWarehouses(alias); err == nil || hasTotalCount {
		pager, err := newWarehousePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &WarehouseConnection{Edges: []*WarehouseEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return v.QueryWarehouses().Paginate(ctx, after, first, before, last, opts...)
}

func (v *Vendor) Products(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ProductOrder, where *ProductWhereInput,
) (*ProductConnection, error) {
	opts := []ProductPaginateOption{
		WithProductOrder(orderBy),
		WithProductFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := v.Edges.totalCount[1][alias]
	if nodes, err := v.NamedProducts(alias); err == nil || hasTotalCount {
		pager, err := newProductPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ProductConnection{Edges: []*ProductEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return v.QueryProducts().Paginate(ctx, after, first, before, last, opts...)
}

func (w *Warehouse) Products(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ProductOrder, where *ProductWhereInput,
) (*ProductConnection, error) {
	opts := []ProductPaginateOption{
		WithProductOrder(orderBy),
		WithProductFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := w.Edges.totalCount[0][alias]
	if nodes, err := w.NamedProducts(alias); err == nil || hasTotalCount {
		pager, err := newProductPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ProductConnection{Edges: []*ProductEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return w.QueryProducts().Paginate(ctx, after, first, before, last, opts...)
}

func (w *Warehouse) Vendor(ctx context.Context) (*Vendor, error) {
	result, err := w.Edges.VendorOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryVendor().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Website) Company(ctx context.Context) (*Company, error) {
	result, err := w.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Website) Country(ctx context.Context) (*Country, error) {
	result, err := w.Edges.CountryOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryCountry().Only(ctx)
	}
	return result, MaskNotFound(err)
}
