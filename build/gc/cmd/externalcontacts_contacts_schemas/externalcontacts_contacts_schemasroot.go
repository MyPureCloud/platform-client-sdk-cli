package externalcontacts_contacts_schemas

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_contacts_schemas_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_contacts_schemas_coretypes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_contacts_schemas_limits"
)

func init() {
	externalcontacts_contacts_schemasCmd.AddCommand(externalcontacts_contacts_schemas_versions.Cmdexternalcontacts_contacts_schemas_versions())
	externalcontacts_contacts_schemasCmd.AddCommand(externalcontacts_contacts_schemas_coretypes.Cmdexternalcontacts_contacts_schemas_coretypes())
	externalcontacts_contacts_schemasCmd.AddCommand(externalcontacts_contacts_schemas_limits.Cmdexternalcontacts_contacts_schemas_limits())
	externalcontacts_contacts_schemasCmd.Short = utils.GenerateCustomDescription(externalcontacts_contacts_schemasCmd.Short, externalcontacts_contacts_schemas_versions.Description, externalcontacts_contacts_schemas_coretypes.Description, externalcontacts_contacts_schemas_limits.Description, )
	externalcontacts_contacts_schemasCmd.Long = externalcontacts_contacts_schemasCmd.Short
}
