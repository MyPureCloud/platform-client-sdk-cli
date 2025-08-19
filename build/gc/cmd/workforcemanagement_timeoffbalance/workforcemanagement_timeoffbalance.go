package workforcemanagement_timeoffbalance

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_timeoffbalance", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/timeoffbalance")
	workforcemanagement_timeoffbalanceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_timeoffbalance"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_timeoffbalanceCmd)
}

func Cmdworkforcemanagement_timeoffbalance() *cobra.Command {
	return workforcemanagement_timeoffbalanceCmd
}
