package workforcemanagement_timeoffrequests_evaluate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_timeoffrequests_evaluate", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/timeoffrequests/evaluate")
	workforcemanagement_timeoffrequests_evaluateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_timeoffrequests_evaluate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_timeoffrequests_evaluateCmd)
}

func Cmdworkforcemanagement_timeoffrequests_evaluate() *cobra.Command {
	return workforcemanagement_timeoffrequests_evaluateCmd
}
