package webdeployments_token

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("webdeployments_token", "SWAGGER_OVERRIDE_/api/v2/webdeployments/token")
	webdeployments_tokenCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webdeployments_token"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webdeployments_tokenCmd)
}

func Cmdwebdeployments_token() *cobra.Command {
	return webdeployments_tokenCmd
}
