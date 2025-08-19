package workforcemanagement_historicaldata

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_historicaldata", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/historicaldata")
	workforcemanagement_historicaldataCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_historicaldata"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_historicaldataCmd)
}

func Cmdworkforcemanagement_historicaldata() *cobra.Command {
	return workforcemanagement_historicaldataCmd
}
