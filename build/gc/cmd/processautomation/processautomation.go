package processautomation

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("processautomation", "SWAGGER_OVERRIDE_/api/v2/processautomation")
	processautomationCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("processautomation"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(processautomationCmd)
}

func Cmdprocessautomation() *cobra.Command {
	return processautomationCmd
}
