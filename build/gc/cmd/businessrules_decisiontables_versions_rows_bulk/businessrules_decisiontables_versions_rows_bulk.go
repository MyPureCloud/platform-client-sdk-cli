package businessrules_decisiontables_versions_rows_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("businessrules_decisiontables_versions_rows_bulk", "SWAGGER_OVERRIDE_/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows/bulk")
	businessrules_decisiontables_versions_rows_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("businessrules_decisiontables_versions_rows_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(businessrules_decisiontables_versions_rows_bulkCmd)
}

func Cmdbusinessrules_decisiontables_versions_rows_bulk() *cobra.Command {
	return businessrules_decisiontables_versions_rows_bulkCmd
}
