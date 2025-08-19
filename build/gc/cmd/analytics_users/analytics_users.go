package analytics_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_users", "SWAGGER_OVERRIDE_/api/v2/analytics/users")
	analytics_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_usersCmd)
}

func Cmdanalytics_users() *cobra.Command {
	return analytics_usersCmd
}
