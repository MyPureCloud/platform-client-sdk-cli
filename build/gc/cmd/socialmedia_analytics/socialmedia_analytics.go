package socialmedia_analytics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("socialmedia_analytics", "SWAGGER_OVERRIDE_/api/v2/socialmedia/analytics")
	socialmedia_analyticsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("socialmedia_analytics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(socialmedia_analyticsCmd)
}

func Cmdsocialmedia_analytics() *cobra.Command {
	return socialmedia_analyticsCmd
}
