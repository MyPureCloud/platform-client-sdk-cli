package flows_validate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("flows_validate", "SWAGGER_OVERRIDE_/api/v2/flows/validate")
	flows_validateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_validate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_validateCmd)
}

func Cmdflows_validate() *cobra.Command {
	return flows_validateCmd
}
