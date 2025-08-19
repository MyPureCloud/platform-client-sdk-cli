package oauth

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("oauth", "SWAGGER_OVERRIDE_/api/v2/oauth")
	oauthCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("oauth"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(oauthCmd)
}

func Cmdoauth() *cobra.Command {
	return oauthCmd
}
