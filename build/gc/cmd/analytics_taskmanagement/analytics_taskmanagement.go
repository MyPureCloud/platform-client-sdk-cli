package analytics_taskmanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_taskmanagement", "SWAGGER_OVERRIDE_/api/v2/analytics/taskmanagement")
	analytics_taskmanagementCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_taskmanagement"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_taskmanagementCmd)
}

func Cmdanalytics_taskmanagement() *cobra.Command {
	return analytics_taskmanagementCmd
}
