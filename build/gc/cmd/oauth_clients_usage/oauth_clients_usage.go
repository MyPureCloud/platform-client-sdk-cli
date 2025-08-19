package oauth_clients_usage

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("oauth_clients_usage", "SWAGGER_OVERRIDE_/api/v2/oauth/clients/{clientId}/usage")
	oauth_clients_usageCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("oauth_clients_usage"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(oauth_clients_usageCmd)
}

func Cmdoauth_clients_usage() *cobra.Command {
	return oauth_clients_usageCmd
}
