package analytics_casemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_casemanagement", "SWAGGER_OVERRIDE_/api/v2/analytics/casemanagement")
	analytics_casemanagementCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_casemanagement"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_casemanagementCmd)
}

func Cmdanalytics_casemanagement() *cobra.Command {
	return analytics_casemanagementCmd
}
