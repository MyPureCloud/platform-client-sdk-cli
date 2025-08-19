package taskmanagement_workitems_bulk_add

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("taskmanagement_workitems_bulk_add", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/workitems/bulk/add")
	taskmanagement_workitems_bulk_addCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_workitems_bulk_add"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_workitems_bulk_addCmd)
}

func Cmdtaskmanagement_workitems_bulk_add() *cobra.Command {
	return taskmanagement_workitems_bulk_addCmd
}
