package taskmanagement_workitems_bulk_terminate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("taskmanagement_workitems_bulk_terminate", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/workitems/bulk/terminate")
	taskmanagement_workitems_bulk_terminateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_workitems_bulk_terminate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_workitems_bulk_terminateCmd)
}

func Cmdtaskmanagement_workitems_bulk_terminate() *cobra.Command {
	return taskmanagement_workitems_bulk_terminateCmd
}
