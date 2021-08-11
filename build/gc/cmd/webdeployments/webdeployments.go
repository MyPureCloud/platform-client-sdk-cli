package webdeployments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("webdeployments", "SWAGGER_OVERRIDE_/api/v2/webdeployments")
	webdeploymentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webdeployments"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webdeploymentsCmd)
}

func Cmdwebdeployments() *cobra.Command {
	return webdeploymentsCmd
}
