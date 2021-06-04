package contentmanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("contentmanagement", "SWAGGER_OVERRIDE_/api/v2/contentmanagement")
	contentmanagementCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("contentmanagement"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(contentmanagementCmd)
}

func Cmdcontentmanagement() *cobra.Command {
	return contentmanagementCmd
}
