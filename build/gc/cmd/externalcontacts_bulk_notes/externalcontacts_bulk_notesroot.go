package externalcontacts_bulk_notes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk_notes_add"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk_notes_update"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk_notes_remove"
)

func init() {
	externalcontacts_bulk_notesCmd.AddCommand(externalcontacts_bulk_notes_add.Cmdexternalcontacts_bulk_notes_add())
	externalcontacts_bulk_notesCmd.AddCommand(externalcontacts_bulk_notes_update.Cmdexternalcontacts_bulk_notes_update())
	externalcontacts_bulk_notesCmd.AddCommand(externalcontacts_bulk_notes_remove.Cmdexternalcontacts_bulk_notes_remove())
	externalcontacts_bulk_notesCmd.Short = utils.GenerateCustomDescription(externalcontacts_bulk_notesCmd.Short, externalcontacts_bulk_notes_add.Description, externalcontacts_bulk_notes_update.Description, externalcontacts_bulk_notes_remove.Description, )
	externalcontacts_bulk_notesCmd.Long = externalcontacts_bulk_notesCmd.Short
}
