package externalcontacts_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk_organizations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk_contacts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk_notes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk_relationships"
)

func init() {
	externalcontacts_bulkCmd.AddCommand(externalcontacts_bulk_organizations.Cmdexternalcontacts_bulk_organizations())
	externalcontacts_bulkCmd.AddCommand(externalcontacts_bulk_contacts.Cmdexternalcontacts_bulk_contacts())
	externalcontacts_bulkCmd.AddCommand(externalcontacts_bulk_notes.Cmdexternalcontacts_bulk_notes())
	externalcontacts_bulkCmd.AddCommand(externalcontacts_bulk_relationships.Cmdexternalcontacts_bulk_relationships())
	externalcontacts_bulkCmd.Short = utils.GenerateCustomDescription(externalcontacts_bulkCmd.Short, externalcontacts_bulk_organizations.Description, externalcontacts_bulk_contacts.Description, externalcontacts_bulk_notes.Description, externalcontacts_bulk_relationships.Description, )
	externalcontacts_bulkCmd.Long = externalcontacts_bulkCmd.Short
}
