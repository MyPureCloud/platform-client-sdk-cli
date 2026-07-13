package externalcontacts_notes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_notes_exports"
)

func init() {
	externalcontacts_notesCmd.AddCommand(externalcontacts_notes_exports.Cmdexternalcontacts_notes_exports())
	externalcontacts_notesCmd.Short = utils.GenerateCustomDescription(externalcontacts_notesCmd.Short, externalcontacts_notes_exports.Description, )
	externalcontacts_notesCmd.Long = externalcontacts_notesCmd.Short
}
