package socialmedia

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("socialmedia", "SWAGGER_OVERRIDE_/api/v2/socialmedia")
	socialmediaCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("socialmedia"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(socialmediaCmd)
}

func Cmdsocialmedia() *cobra.Command {
	return socialmediaCmd
}
