package workforcemanagement_selfschedule_offers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_selfschedule_offers", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/selfschedule/offers")
	workforcemanagement_selfschedule_offersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_selfschedule_offers"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_selfschedule_offersCmd)
}

func Cmdworkforcemanagement_selfschedule_offers() *cobra.Command {
	return workforcemanagement_selfschedule_offersCmd
}
