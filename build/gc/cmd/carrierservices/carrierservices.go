package carrierservices

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("carrierservices", "SWAGGER_OVERRIDE_/api/v2/carrierservices")
	carrierservicesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("carrierservices"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(carrierservicesCmd)
}

func Cmdcarrierservices() *cobra.Command {
	return carrierservicesCmd
}
