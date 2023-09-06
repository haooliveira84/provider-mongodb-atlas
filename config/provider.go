/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/haooliveira84/provider-mongodbatlas/config/backup"
	"github.com/haooliveira84/provider-mongodbatlas/config/cluster"
	"github.com/haooliveira84/provider-mongodbatlas/config/dbrole"
	"github.com/haooliveira84/provider-mongodbatlas/config/dbuser"
	"github.com/haooliveira84/provider-mongodbatlas/config/organization"
	"github.com/haooliveira84/provider-mongodbatlas/config/project"
	"github.com/haooliveira84/provider-mongodbatlas/config/teams"
)

const (
	resourcePrefix = "mongodbatlas"
	modulePath     = "github.com/haooliveira84/provider-mongodbatlas"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("mongodbatlas.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		backup.Configure,
		cluster.Configure,
		dbrole.Configure,
		dbuser.Configure,
		organization.Configure,
		project.Configure,
		teams.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
