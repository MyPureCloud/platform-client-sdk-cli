package journey

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("journey", "SWAGGER_OVERRIDE_/api/v2/journey")
	journeyCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journeyCmd)
}

func Cmdjourney() *cobra.Command {
	return journeyCmd
}
