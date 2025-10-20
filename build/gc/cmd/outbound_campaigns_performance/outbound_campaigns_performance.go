package outbound_campaigns_performance

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("outbound_campaigns_performance", "SWAGGER_OVERRIDE_/api/v2/outbound/campaigns/performance")
	outbound_campaigns_performanceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_campaigns_performance"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_campaigns_performanceCmd)
}

func Cmdoutbound_campaigns_performance() *cobra.Command {
	return outbound_campaigns_performanceCmd
}
