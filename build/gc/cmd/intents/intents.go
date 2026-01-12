package intents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("intents", "SWAGGER_OVERRIDE_/api/v2/intents")
	intentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("intents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(intentsCmd)
}

func Cmdintents() *cobra.Command {
	return intentsCmd
}
