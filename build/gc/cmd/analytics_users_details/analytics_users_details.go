package analytics_users_details

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_users_details", "SWAGGER_OVERRIDE_/api/v2/analytics/users/details")
	analytics_users_detailsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_users_details"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_users_detailsCmd)
}

func Cmdanalytics_users_details() *cobra.Command {
	return analytics_users_detailsCmd
}
