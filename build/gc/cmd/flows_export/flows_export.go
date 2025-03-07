package flows_export

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("flows_export", "SWAGGER_OVERRIDE_/api/v2/flows/export")
	flows_exportCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_export"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_exportCmd)
}

func Cmdflows_export() *cobra.Command {
	return flows_exportCmd
}
