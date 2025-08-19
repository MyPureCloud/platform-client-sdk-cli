package outbound_campaigns_callback

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("outbound_campaigns_callback", "SWAGGER_OVERRIDE_/api/v2/outbound/campaigns/{campaignId}/callback")
	outbound_campaigns_callbackCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_campaigns_callback"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_campaigns_callbackCmd)
}

func Cmdoutbound_campaigns_callback() *cobra.Command {
	return outbound_campaigns_callbackCmd
}
