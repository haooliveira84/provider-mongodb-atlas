/*
Copyright 2021 Upbound Inc.
*/

package organization

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("mongodbatlas_organization", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.ShortGroup = ""
	})
	p.AddResourceConfigurator("mongodbatlas_organizations", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.ShortGroup = ""
	})
}
