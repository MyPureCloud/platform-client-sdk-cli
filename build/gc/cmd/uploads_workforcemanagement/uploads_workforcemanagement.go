package uploads_workforcemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("uploads_workforcemanagement", "SWAGGER_OVERRIDE_/api/v2/uploads/workforcemanagement")
	uploads_workforcemanagementCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("uploads_workforcemanagement"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(uploads_workforcemanagementCmd)
}

func Cmduploads_workforcemanagement() *cobra.Command {
	return uploads_workforcemanagementCmd
}
