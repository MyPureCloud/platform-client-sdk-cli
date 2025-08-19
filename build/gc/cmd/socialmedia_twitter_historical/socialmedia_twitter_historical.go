package socialmedia_twitter_historical

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("socialmedia_twitter_historical", "SWAGGER_OVERRIDE_/api/v2/socialmedia/twitter/historical")
	socialmedia_twitter_historicalCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("socialmedia_twitter_historical"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(socialmedia_twitter_historicalCmd)
}

func Cmdsocialmedia_twitter_historical() *cobra.Command {
	return socialmedia_twitter_historicalCmd
}
