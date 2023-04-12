package entproject

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"

	"github.com/entkit/entkit"
	"github.com/entkit/entkit-demo/ent"
	"github.com/google/uuid"
)

// CreateVendor is the resolver for the createVendor field.
func (r *mutationResolver) CreateVendor(ctx context.Context, input ent.CreateVendorInput) (*ent.Vendor, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Vendor", entkit.ActionCreateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Vendor.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateVendor is the resolver for the updateVendor field.
func (r *mutationResolver) UpdateVendor(ctx context.Context, id uuid.UUID, input ent.UpdateVendorInput) (*ent.Vendor, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Vendor", entkit.ActionUpdateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Vendor.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteVendors is the resolver for the deleteVendors field.
func (r *mutationResolver) DeleteVendors(ctx context.Context, where ent.VendorWhereInput) (int, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Vendor", entkit.ActionDeleteScope)
	if err != nil {
		return 0, err
	}
	p, err := where.P()
	if err != nil {
		return 0, err
	}
	return ent.FromContext(ctx).Vendor.
		Delete().
		Where(p).
		Exec(ctx)
}

// CreateWarehouse is the resolver for the createWarehouse field.
func (r *mutationResolver) CreateWarehouse(ctx context.Context, input ent.CreateWarehouseInput) (*ent.Warehouse, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Warehouse", entkit.ActionCreateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Warehouse.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateWarehouse is the resolver for the updateWarehouse field.
func (r *mutationResolver) UpdateWarehouse(ctx context.Context, id uuid.UUID, input ent.UpdateWarehouseInput) (*ent.Warehouse, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Warehouse", entkit.ActionUpdateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Warehouse.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteWarehouses is the resolver for the deleteWarehouses field.
func (r *mutationResolver) DeleteWarehouses(ctx context.Context, where ent.WarehouseWhereInput) (int, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Warehouse", entkit.ActionDeleteScope)
	if err != nil {
		return 0, err
	}
	p, err := where.P()
	if err != nil {
		return 0, err
	}
	return ent.FromContext(ctx).Warehouse.
		Delete().
		Where(p).
		Exec(ctx)
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input ent.CreateProductInput) (*ent.Product, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Product", entkit.ActionCreateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Product.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, id uuid.UUID, input ent.UpdateProductInput) (*ent.Product, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Product", entkit.ActionUpdateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Product.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteProducts is the resolver for the deleteProducts field.
func (r *mutationResolver) DeleteProducts(ctx context.Context, where ent.ProductWhereInput) (int, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Product", entkit.ActionDeleteScope)
	if err != nil {
		return 1, err
	}
	p, err := where.P()
	if err != nil {
		return 0, err
	}
	return ent.FromContext(ctx).Product.
		Delete().
		Where(p).
		Exec(ctx)
}

// CreateCompany is the resolver for the createCompany field.
func (r *mutationResolver) CreateCompany(ctx context.Context, input ent.CreateCompanyInput) (*ent.Company, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Company", entkit.ActionCreateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Company.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateCompany is the resolver for the updateCompany field.
func (r *mutationResolver) UpdateCompany(ctx context.Context, id uuid.UUID, input ent.UpdateCompanyInput) (*ent.Company, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Company", entkit.ActionUpdateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Company.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteCompanies is the resolver for the deleteCompanies field.
func (r *mutationResolver) DeleteCompanies(ctx context.Context, where ent.CompanyWhereInput) (int, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Company", entkit.ActionDeleteScope)
	if err != nil {
		return 0, err
	}
	p, err := where.P()
	if err != nil {
		return 0, err
	}
	return ent.FromContext(ctx).Company.
		Delete().
		Where(p).
		Exec(ctx)
}

// CreateCountry is the resolver for the createCountry field.
func (r *mutationResolver) CreateCountry(ctx context.Context, input ent.CreateCountryInput) (*ent.Country, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Country", entkit.ActionCreateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Country.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateCountry is the resolver for the updateCountry field.
func (r *mutationResolver) UpdateCountry(ctx context.Context, id uuid.UUID, input ent.UpdateCountryInput) (*ent.Country, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Country", entkit.ActionUpdateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Country.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteCountries is the resolver for the deleteCountries field.
func (r *mutationResolver) DeleteCountries(ctx context.Context, where ent.CountryWhereInput) (int, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Country", entkit.ActionDeleteScope)
	if err != nil {
		return 0, err
	}
	p, err := where.P()
	if err != nil {
		return 0, err
	}
	return ent.FromContext(ctx).Country.
		Delete().
		Where(p).
		Exec(ctx)
}

// CreatePhone is the resolver for the createPhone field.
func (r *mutationResolver) CreatePhone(ctx context.Context, input ent.CreatePhoneInput) (*ent.Phone, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Phone", entkit.ActionCreateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Phone.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdatePhone is the resolver for the updatePhone field.
func (r *mutationResolver) UpdatePhone(ctx context.Context, id uuid.UUID, input ent.UpdatePhoneInput) (*ent.Phone, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Phone", entkit.ActionUpdateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Phone.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeletePhones is the resolver for the deletePhones field.
func (r *mutationResolver) DeletePhones(ctx context.Context, where ent.PhoneWhereInput) (int, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Phone", entkit.ActionDeleteScope)
	if err != nil {
		return 0, err
	}
	p, err := where.P()
	if err != nil {
		return 0, err
	}
	return ent.FromContext(ctx).Phone.
		Delete().
		Where(p).
		Exec(ctx)
}

// CreateLocation is the resolver for the createLocation field.
func (r *mutationResolver) CreateLocation(ctx context.Context, input ent.CreateLocationInput) (*ent.Location, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Location", entkit.ActionCreateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Location.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateLocation is the resolver for the updateLocation field.
func (r *mutationResolver) UpdateLocation(ctx context.Context, id uuid.UUID, input ent.UpdateLocationInput) (*ent.Location, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Location", entkit.ActionUpdateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Location.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteLocations is the resolver for the deleteLocations field.
func (r *mutationResolver) DeleteLocations(ctx context.Context, where ent.LocationWhereInput) (int, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Location", entkit.ActionDeleteScope)
	if err != nil {
		return 0, err
	}
	p, err := where.P()
	if err != nil {
		return 0, err
	}
	return ent.FromContext(ctx).Location.
		Delete().
		Where(p).
		Exec(ctx)
}

// CreateWebsite is the resolver for the createWebsite field.
func (r *mutationResolver) CreateWebsite(ctx context.Context, input ent.CreateWebsiteInput) (*ent.Website, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Website", entkit.ActionCreateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Website.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateWebsite is the resolver for the updateWebsite field.
func (r *mutationResolver) UpdateWebsite(ctx context.Context, id uuid.UUID, input ent.UpdateWebsiteInput) (*ent.Website, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Website", entkit.ActionUpdateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Website.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteWebsites is the resolver for the deleteWebsites field.
func (r *mutationResolver) DeleteWebsites(ctx context.Context, where ent.WebsiteWhereInput) (int, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Website", entkit.ActionDeleteScope)
	if err != nil {
		return 0, err
	}
	p, err := where.P()
	if err != nil {
		return 0, err
	}
	return ent.FromContext(ctx).Website.
		Delete().
		Where(p).
		Exec(ctx)
}

// CreateEmail is the resolver for the createEmail field.
func (r *mutationResolver) CreateEmail(ctx context.Context, input ent.CreateEmailInput) (*ent.Email, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Email", entkit.ActionCreateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Email.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateEmail is the resolver for the updateEmail field.
func (r *mutationResolver) UpdateEmail(ctx context.Context, id uuid.UUID, input ent.UpdateEmailInput) (*ent.Email, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Email", entkit.ActionUpdateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Email.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteEmails is the resolver for the deleteEmails field.
func (r *mutationResolver) DeleteEmails(ctx context.Context, where ent.EmailWhereInput) (int, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Email", entkit.ActionDeleteScope)
	if err != nil {
		return 0, err
	}
	p, err := where.P()
	if err != nil {
		return 0, err
	}
	return ent.FromContext(ctx).Email.
		Delete().
		Where(p).
		Exec(ctx)
}

// CreateImage is the resolver for the createImage field.
func (r *mutationResolver) CreateImage(ctx context.Context, input ent.CreateImageInput) (*ent.Image, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Image", entkit.ActionCreateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Image.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateImage is the resolver for the updateImage field.
func (r *mutationResolver) UpdateImage(ctx context.Context, id uuid.UUID, input ent.UpdateImageInput) (*ent.Image, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Image", entkit.ActionUpdateScope)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).Image.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteImages is the resolver for the deleteImages field.
func (r *mutationResolver) DeleteImages(ctx context.Context, where ent.ImageWhereInput) (int, error) {
	err := ent.EntkitAuthorizeByResource(ctx, "Image", entkit.ActionDeleteScope)
	if err != nil {
		return 0, err
	}
	p, err := where.P()
	if err != nil {
		return 0, err
	}
	return ent.FromContext(ctx).Image.
		Delete().
		Where(p).
		Exec(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
