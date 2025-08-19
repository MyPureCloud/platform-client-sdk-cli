package workforcemanagement_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_users", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/users")
	workforcemanagement_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_usersCmd)
}

func Cmdworkforcemanagement_users() *cobra.Command {
	return workforcemanagement_usersCmd
}
