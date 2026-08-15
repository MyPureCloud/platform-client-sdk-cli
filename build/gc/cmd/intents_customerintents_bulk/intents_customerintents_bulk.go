package intents_customerintents_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("intents_customerintents_bulk", "SWAGGER_OVERRIDE_/api/v2/intents/customerintents/bulk")
	intents_customerintents_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("intents_customerintents_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(intents_customerintents_bulkCmd)
}

func Cmdintents_customerintents_bulk() *cobra.Command {
	return intents_customerintents_bulkCmd
}
