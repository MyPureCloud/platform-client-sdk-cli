package responsemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("responsemanagement", "SWAGGER_OVERRIDE_/api/v2/responsemanagement")
	responsemanagementCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("responsemanagement"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(responsemanagementCmd)
}

func Cmdresponsemanagement() *cobra.Command {
	return responsemanagementCmd
}
