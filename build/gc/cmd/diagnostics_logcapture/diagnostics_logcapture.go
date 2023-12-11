package diagnostics_logcapture

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("diagnostics_logcapture", "SWAGGER_OVERRIDE_/api/v2/diagnostics/logcapture")
	diagnostics_logcaptureCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("diagnostics_logcapture"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(diagnostics_logcaptureCmd)
}

func Cmddiagnostics_logcapture() *cobra.Command {
	return diagnostics_logcaptureCmd
}
