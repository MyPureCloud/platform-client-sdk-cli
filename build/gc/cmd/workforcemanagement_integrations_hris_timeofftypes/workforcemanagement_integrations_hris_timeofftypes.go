package workforcemanagement_integrations_hris_timeofftypes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_integrations_hris_timeofftypes", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/integrations/hris/timeofftypes")
	workforcemanagement_integrations_hris_timeofftypesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_integrations_hris_timeofftypes"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_integrations_hris_timeofftypesCmd)
}

func Cmdworkforcemanagement_integrations_hris_timeofftypes() *cobra.Command {
	return workforcemanagement_integrations_hris_timeofftypesCmd
}
