package analytics_reporting_settings_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_reporting_settings_users", "SWAGGER_OVERRIDE_/api/v2/analytics/reporting/settings/users")
	analytics_reporting_settings_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_reporting_settings_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_reporting_settings_usersCmd)
}

func Cmdanalytics_reporting_settings_users() *cobra.Command {
	return analytics_reporting_settings_usersCmd
}
