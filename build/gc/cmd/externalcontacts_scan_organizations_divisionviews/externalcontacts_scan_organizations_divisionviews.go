package externalcontacts_scan_organizations_divisionviews

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("externalcontacts_scan_organizations_divisionviews", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/scan/organizations/divisionviews")
	externalcontacts_scan_organizations_divisionviewsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_scan_organizations_divisionviews"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_scan_organizations_divisionviewsCmd)
}

func Cmdexternalcontacts_scan_organizations_divisionviews() *cobra.Command {
	return externalcontacts_scan_organizations_divisionviewsCmd
}
