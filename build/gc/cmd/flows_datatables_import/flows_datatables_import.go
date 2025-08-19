package flows_datatables_import

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("flows_datatables_import", "SWAGGER_OVERRIDE_/api/v2/flows/datatables/{datatableId}/import")
	flows_datatables_importCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_datatables_import"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_datatables_importCmd)
}

func Cmdflows_datatables_import() *cobra.Command {
	return flows_datatables_importCmd
}
