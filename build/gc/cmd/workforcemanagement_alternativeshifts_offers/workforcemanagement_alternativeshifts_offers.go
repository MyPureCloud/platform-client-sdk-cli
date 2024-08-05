package workforcemanagement_alternativeshifts_offers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_alternativeshifts_offers", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/alternativeshifts/offers")
	workforcemanagement_alternativeshifts_offersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_alternativeshifts_offers"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_alternativeshifts_offersCmd)
}

func Cmdworkforcemanagement_alternativeshifts_offers() *cobra.Command {
	return workforcemanagement_alternativeshifts_offersCmd
}
