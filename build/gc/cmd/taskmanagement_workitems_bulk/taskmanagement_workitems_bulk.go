package taskmanagement_workitems_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("taskmanagement_workitems_bulk", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/workitems/bulk")
	taskmanagement_workitems_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_workitems_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_workitems_bulkCmd)
}

func Cmdtaskmanagement_workitems_bulk() *cobra.Command {
	return taskmanagement_workitems_bulkCmd
}
