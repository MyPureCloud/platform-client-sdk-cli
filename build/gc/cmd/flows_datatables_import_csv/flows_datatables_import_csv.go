package flows_datatables_import_csv

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("flows_datatables_import_csv", "SWAGGER_OVERRIDE_/api/v2/flows/datatables/{datatableId}/import/csv")
	flows_datatables_import_csvCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_datatables_import_csv"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_datatables_import_csvCmd)
}

func Cmdflows_datatables_import_csv() *cobra.Command {
	return flows_datatables_import_csvCmd
}
