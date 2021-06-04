package outbound

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("outbound", "SWAGGER_OVERRIDE_/api/v2/outbound")
	outboundCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outboundCmd)
}

func Cmdoutbound() *cobra.Command {
	return outboundCmd
}
