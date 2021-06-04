package authorization

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("authorization", "SWAGGER_OVERRIDE_/api/v2/authorization")
	authorizationCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorizationCmd)
}

func Cmdauthorization() *cobra.Command {
	return authorizationCmd
}
