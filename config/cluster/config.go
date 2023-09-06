/*
Copyright 2021 Upbound Inc.
*/

package cluster

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("mongodbatlas_cluster", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.ShortGroup = ""
	})
}
