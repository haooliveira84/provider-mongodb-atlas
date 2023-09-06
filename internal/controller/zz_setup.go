/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	listapikey "github.com/haooliveira84/provider-mongodbatlas/internal/controller/access/listapikey"
	cluster "github.com/haooliveira84/provider-mongodbatlas/internal/controller/advanced/cluster"
	configuration "github.com/haooliveira84/provider-mongodbatlas/internal/controller/alert/configuration"
	key "github.com/haooliveira84/provider-mongodbatlas/internal/controller/api/key"
	compliancepolicy "github.com/haooliveira84/provider-mongodbatlas/internal/controller/backup/compliancepolicy"
	backupschedule "github.com/haooliveira84/provider-mongodbatlas/internal/controller/cloud/backupschedule"
	backupsnapshotexportbucket "github.com/haooliveira84/provider-mongodbatlas/internal/controller/cloud/backupsnapshotexportbucket"
	backupsnapshotexportjob "github.com/haooliveira84/provider-mongodbatlas/internal/controller/cloud/backupsnapshotexportjob"
	backupsnapshotrestorejob "github.com/haooliveira84/provider-mongodbatlas/internal/controller/cloud/backupsnapshotrestorejob"
	provideraccess "github.com/haooliveira84/provider-mongodbatlas/internal/controller/cloud/provideraccess"
	provideraccessauthorization "github.com/haooliveira84/provider-mongodbatlas/internal/controller/cloud/provideraccessauthorization"
	provideraccesssetup "github.com/haooliveira84/provider-mongodbatlas/internal/controller/cloud/provideraccesssetup"
	clustercluster "github.com/haooliveira84/provider-mongodbatlas/internal/controller/cluster/cluster"
	outagesimulation "github.com/haooliveira84/provider-mongodbatlas/internal/controller/cluster/outagesimulation"
	dnsconfigurationclusteraws "github.com/haooliveira84/provider-mongodbatlas/internal/controller/custom/dnsconfigurationclusteraws"
	lake "github.com/haooliveira84/provider-mongodbatlas/internal/controller/data/lake"
	lakepipeline "github.com/haooliveira84/provider-mongodbatlas/internal/controller/data/lakepipeline"
	atrest "github.com/haooliveira84/provider-mongodbatlas/internal/controller/encryption/atrest"
	trigger "github.com/haooliveira84/provider-mongodbatlas/internal/controller/event/trigger"
	databaseinstance "github.com/haooliveira84/provider-mongodbatlas/internal/controller/federated/databaseinstance"
	querylimit "github.com/haooliveira84/provider-mongodbatlas/internal/controller/federated/querylimit"
	settingsidentityprovider "github.com/haooliveira84/provider-mongodbatlas/internal/controller/federated/settingsidentityprovider"
	settingsorgconfig "github.com/haooliveira84/provider-mongodbatlas/internal/controller/federated/settingsorgconfig"
	settingsorgrolemapping "github.com/haooliveira84/provider-mongodbatlas/internal/controller/federated/settingsorgrolemapping"
	clusterconfig "github.com/haooliveira84/provider-mongodbatlas/internal/controller/global/clusterconfig"
	configurationldap "github.com/haooliveira84/provider-mongodbatlas/internal/controller/ldap/configuration"
	verify "github.com/haooliveira84/provider-mongodbatlas/internal/controller/ldap/verify"
	window "github.com/haooliveira84/provider-mongodbatlas/internal/controller/maintenance/window"
	auditing "github.com/haooliveira84/provider-mongodbatlas/internal/controller/mongodbatlas/auditing"
	dbrole "github.com/haooliveira84/provider-mongodbatlas/internal/controller/mongodbatlas/dbrole"
	organization "github.com/haooliveira84/provider-mongodbatlas/internal/controller/mongodbatlas/organization"
	project "github.com/haooliveira84/provider-mongodbatlas/internal/controller/mongodbatlas/project"
	team "github.com/haooliveira84/provider-mongodbatlas/internal/controller/mongodbatlas/team"
	teams "github.com/haooliveira84/provider-mongodbatlas/internal/controller/mongodbatlas/teams"
	user "github.com/haooliveira84/provider-mongodbatlas/internal/controller/mongodbatlas/user"
	container "github.com/haooliveira84/provider-mongodbatlas/internal/controller/network/container"
	peering "github.com/haooliveira84/provider-mongodbatlas/internal/controller/network/peering"
	archive "github.com/haooliveira84/provider-mongodbatlas/internal/controller/online/archive"
	invitation "github.com/haooliveira84/provider-mongodbatlas/internal/controller/org/invitation"
	endpointregionalmode "github.com/haooliveira84/provider-mongodbatlas/internal/controller/private/endpointregionalmode"
	endpoint "github.com/haooliveira84/provider-mongodbatlas/internal/controller/privatelink/endpoint"
	endpointserverless "github.com/haooliveira84/provider-mongodbatlas/internal/controller/privatelink/endpointserverless"
	endpointservice "github.com/haooliveira84/provider-mongodbatlas/internal/controller/privatelink/endpointservice"
	endpointserviceadl "github.com/haooliveira84/provider-mongodbatlas/internal/controller/privatelink/endpointserviceadl"
	endpointservicedatafederationonlinearchive "github.com/haooliveira84/provider-mongodbatlas/internal/controller/privatelink/endpointservicedatafederationonlinearchive"
	endpointserviceserverless "github.com/haooliveira84/provider-mongodbatlas/internal/controller/privatelink/endpointserviceserverless"
	apikey "github.com/haooliveira84/provider-mongodbatlas/internal/controller/project/apikey"
	invitationproject "github.com/haooliveira84/provider-mongodbatlas/internal/controller/project/invitation"
	ipaccesslist "github.com/haooliveira84/provider-mongodbatlas/internal/controller/project/ipaccesslist"
	providerconfig "github.com/haooliveira84/provider-mongodbatlas/internal/controller/providerconfig"
	index "github.com/haooliveira84/provider-mongodbatlas/internal/controller/search/index"
	instance "github.com/haooliveira84/provider-mongodbatlas/internal/controller/serverless/instance"
	backupsnapshot "github.com/haooliveira84/provider-mongodbatlas/internal/controller/snapshot/backupsnapshot"
	partyintegration "github.com/haooliveira84/provider-mongodbatlas/internal/controller/third/partyintegration"
	authenticationdatabaseuser "github.com/haooliveira84/provider-mongodbatlas/internal/controller/x509/authenticationdatabaseuser"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		listapikey.Setup,
		cluster.Setup,
		configuration.Setup,
		key.Setup,
		compliancepolicy.Setup,
		backupschedule.Setup,
		backupsnapshotexportbucket.Setup,
		backupsnapshotexportjob.Setup,
		backupsnapshotrestorejob.Setup,
		provideraccess.Setup,
		provideraccessauthorization.Setup,
		provideraccesssetup.Setup,
		clustercluster.Setup,
		outagesimulation.Setup,
		dnsconfigurationclusteraws.Setup,
		lake.Setup,
		lakepipeline.Setup,
		atrest.Setup,
		trigger.Setup,
		databaseinstance.Setup,
		querylimit.Setup,
		settingsidentityprovider.Setup,
		settingsorgconfig.Setup,
		settingsorgrolemapping.Setup,
		clusterconfig.Setup,
		configurationldap.Setup,
		verify.Setup,
		window.Setup,
		auditing.Setup,
		dbrole.Setup,
		organization.Setup,
		project.Setup,
		team.Setup,
		teams.Setup,
		user.Setup,
		container.Setup,
		peering.Setup,
		archive.Setup,
		invitation.Setup,
		endpointregionalmode.Setup,
		endpoint.Setup,
		endpointserverless.Setup,
		endpointservice.Setup,
		endpointserviceadl.Setup,
		endpointservicedatafederationonlinearchive.Setup,
		endpointserviceserverless.Setup,
		apikey.Setup,
		invitationproject.Setup,
		ipaccesslist.Setup,
		providerconfig.Setup,
		index.Setup,
		instance.Setup,
		backupsnapshot.Setup,
		partyintegration.Setup,
		authenticationdatabaseuser.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
