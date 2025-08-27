package webmessaging_deployments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("webmessaging_deployments", "SWAGGER_OVERRIDE_/api/v2/webmessaging/deployments")
	webmessaging_deploymentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webmessaging_deployments"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webmessaging_deploymentsCmd)
}

func Cmdwebmessaging_deployments() *cobra.Command {
	return webmessaging_deploymentsCmd
}
