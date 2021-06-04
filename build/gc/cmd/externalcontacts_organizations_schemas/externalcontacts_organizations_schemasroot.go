package externalcontacts_organizations_schemas

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_organizations_schemas_versions"
)

func init() {
	externalcontacts_organizations_schemasCmd.AddCommand(externalcontacts_organizations_schemas_versions.Cmdexternalcontacts_organizations_schemas_versions())
	externalcontacts_organizations_schemasCmd.Short = utils.GenerateCustomDescription(externalcontacts_organizations_schemasCmd.Short, externalcontacts_organizations_schemas_versions.Description, )
	externalcontacts_organizations_schemasCmd.Long = externalcontacts_organizations_schemasCmd.Short
}
