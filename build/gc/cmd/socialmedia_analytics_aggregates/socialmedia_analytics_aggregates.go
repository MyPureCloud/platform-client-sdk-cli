package socialmedia_analytics_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("socialmedia_analytics_aggregates", "SWAGGER_OVERRIDE_/api/v2/socialmedia/analytics/aggregates")
	socialmedia_analytics_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("socialmedia_analytics_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(socialmedia_analytics_aggregatesCmd)
}

func Cmdsocialmedia_analytics_aggregates() *cobra.Command {
	return socialmedia_analytics_aggregatesCmd
}
