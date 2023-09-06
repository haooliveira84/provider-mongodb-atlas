/*
Copyright 2021 Upbound Inc.
*/

package backup

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the Backup group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("mongodbatlas_cloud_backup_snapshot", func(r *ujconfig.Resource) {
		r.ShortGroup = "snapshot"
	})
}
