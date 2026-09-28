package taskmanagement_workitems_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("taskmanagement_workitems_users", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/workitems/{workitemId}/users/me")
	taskmanagement_workitems_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_workitems_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_workitems_usersCmd)
}

func Cmdtaskmanagement_workitems_users() *cobra.Command {
	return taskmanagement_workitems_usersCmd
}
