package taskmanagement_workitems_users_me

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("taskmanagement_workitems_users_me", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/workitems/{workitemId}/users/me")
	taskmanagement_workitems_users_meCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_workitems_users_me"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_workitems_users_meCmd)
}

func Cmdtaskmanagement_workitems_users_me() *cobra.Command {
	return taskmanagement_workitems_users_meCmd
}
