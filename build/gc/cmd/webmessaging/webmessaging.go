package webmessaging

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("webmessaging", "SWAGGER_OVERRIDE_/api/v2/webmessaging")
	webmessagingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webmessaging"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webmessagingCmd)
}

func Cmdwebmessaging() *cobra.Command {
	return webmessagingCmd
}
