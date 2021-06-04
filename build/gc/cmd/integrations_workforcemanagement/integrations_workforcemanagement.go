package integrations_workforcemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_workforcemanagement", "SWAGGER_OVERRIDE_/api/v2/integrations/workforcemanagement")
	integrations_workforcemanagementCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_workforcemanagement"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_workforcemanagementCmd)
}

func Cmdintegrations_workforcemanagement() *cobra.Command {
	return integrations_workforcemanagementCmd
}
