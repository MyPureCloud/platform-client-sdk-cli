package externalcontacts_scan_relationships_divisionviews

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("externalcontacts_scan_relationships_divisionviews", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/scan/relationships/divisionviews")
	externalcontacts_scan_relationships_divisionviewsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_scan_relationships_divisionviews"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_scan_relationships_divisionviewsCmd)
}

func Cmdexternalcontacts_scan_relationships_divisionviews() *cobra.Command {
	return externalcontacts_scan_relationships_divisionviewsCmd
}
