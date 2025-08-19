package externalcontacts_scan_relationships

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_scan_relationships_divisionviews"
)

func init() {
	externalcontacts_scan_relationshipsCmd.AddCommand(externalcontacts_scan_relationships_divisionviews.Cmdexternalcontacts_scan_relationships_divisionviews())
	externalcontacts_scan_relationshipsCmd.Short = utils.GenerateCustomDescription(externalcontacts_scan_relationshipsCmd.Short, externalcontacts_scan_relationships_divisionviews.Description, )
	externalcontacts_scan_relationshipsCmd.Long = externalcontacts_scan_relationshipsCmd.Short
}
