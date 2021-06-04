package uploads_workforcemanagement_historicaldata

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("uploads_workforcemanagement_historicaldata", "SWAGGER_OVERRIDE_/api/v2/uploads/workforcemanagement/historicaldata")
	uploads_workforcemanagement_historicaldataCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("uploads_workforcemanagement_historicaldata"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(uploads_workforcemanagement_historicaldataCmd)
}

func Cmduploads_workforcemanagement_historicaldata() *cobra.Command {
	return uploads_workforcemanagement_historicaldataCmd
}
