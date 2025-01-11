package taskmanagement_workitems_bulk_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("taskmanagement_workitems_bulk_jobs", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/workitems/bulk/jobs")
	taskmanagement_workitems_bulk_jobsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_workitems_bulk_jobs"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_workitems_bulk_jobsCmd)
}

func Cmdtaskmanagement_workitems_bulk_jobs() *cobra.Command {
	return taskmanagement_workitems_bulk_jobsCmd
}
