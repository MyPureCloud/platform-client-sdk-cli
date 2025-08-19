package authorization_divisionspermitted

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("authorization_divisionspermitted", "SWAGGER_OVERRIDE_/api/v2/authorization/divisionspermitted")
	authorization_divisionspermittedCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_divisionspermitted"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_divisionspermittedCmd)
}

func Cmdauthorization_divisionspermitted() *cobra.Command {
	return authorization_divisionspermittedCmd
}
