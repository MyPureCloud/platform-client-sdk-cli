package socialmedia_twitter

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("socialmedia_twitter", "SWAGGER_OVERRIDE_/api/v2/socialmedia/twitter")
	socialmedia_twitterCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("socialmedia_twitter"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(socialmedia_twitterCmd)
}

func Cmdsocialmedia_twitter() *cobra.Command {
	return socialmedia_twitterCmd
}
