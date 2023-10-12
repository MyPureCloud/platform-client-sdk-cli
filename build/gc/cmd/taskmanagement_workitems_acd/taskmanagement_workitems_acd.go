package taskmanagement_workitems_acd

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("taskmanagement_workitems_acd", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/workitems/{workitemId}/acd")
	taskmanagement_workitems_acdCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_workitems_acd"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_workitems_acdCmd)
}

func Cmdtaskmanagement_workitems_acd() *cobra.Command {
	return taskmanagement_workitems_acdCmd
}
