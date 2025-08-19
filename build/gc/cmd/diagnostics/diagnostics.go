package diagnostics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("diagnostics", "SWAGGER_OVERRIDE_/api/v2/diagnostics")
	diagnosticsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("diagnostics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(diagnosticsCmd)
}

func Cmddiagnostics() *cobra.Command {
	return diagnosticsCmd
}
