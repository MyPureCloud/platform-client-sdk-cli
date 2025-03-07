package externalcontacts_scan_organizations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_scan_organizations_divisionviews"
)

func init() {
	externalcontacts_scan_organizationsCmd.AddCommand(externalcontacts_scan_organizations_divisionviews.Cmdexternalcontacts_scan_organizations_divisionviews())
	externalcontacts_scan_organizationsCmd.Short = utils.GenerateCustomDescription(externalcontacts_scan_organizationsCmd.Short, externalcontacts_scan_organizations_divisionviews.Description, )
	externalcontacts_scan_organizationsCmd.Long = externalcontacts_scan_organizationsCmd.Short
}
