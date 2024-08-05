package workforcemanagement_alternativeshifts_offers_search

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_alternativeshifts_offers_search", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/alternativeshifts/offers/search")
	workforcemanagement_alternativeshifts_offers_searchCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_alternativeshifts_offers_search"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_alternativeshifts_offers_searchCmd)
}

func Cmdworkforcemanagement_alternativeshifts_offers_search() *cobra.Command {
	return workforcemanagement_alternativeshifts_offers_searchCmd
}
