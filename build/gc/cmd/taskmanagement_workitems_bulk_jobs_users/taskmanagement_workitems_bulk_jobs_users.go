package taskmanagement_workitems_bulk_jobs_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("taskmanagement_workitems_bulk_jobs_users", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/workitems/bulk/jobs/users")
	taskmanagement_workitems_bulk_jobs_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_workitems_bulk_jobs_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_workitems_bulk_jobs_usersCmd)
}

func Cmdtaskmanagement_workitems_bulk_jobs_users() *cobra.Command {
	return taskmanagement_workitems_bulk_jobs_usersCmd
}
