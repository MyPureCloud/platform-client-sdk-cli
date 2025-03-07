package externalcontacts_scan_notes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_scan_notes_divisionviews"
)

func init() {
	externalcontacts_scan_notesCmd.AddCommand(externalcontacts_scan_notes_divisionviews.Cmdexternalcontacts_scan_notes_divisionviews())
	externalcontacts_scan_notesCmd.Short = utils.GenerateCustomDescription(externalcontacts_scan_notesCmd.Short, externalcontacts_scan_notes_divisionviews.Description, )
	externalcontacts_scan_notesCmd.Long = externalcontacts_scan_notesCmd.Short
}
