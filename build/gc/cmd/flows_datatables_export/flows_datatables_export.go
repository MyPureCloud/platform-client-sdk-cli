package flows_datatables_export

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("flows_datatables_export", "SWAGGER_OVERRIDE_/api/v2/flows/datatables/{datatableId}/export")
	flows_datatables_exportCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_datatables_export"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_datatables_exportCmd)
}

func Cmdflows_datatables_export() *cobra.Command {
	return flows_datatables_exportCmd
}
