package workforcemanagement_adherence_historical

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_adherence_historical", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/adherence/historical")
	workforcemanagement_adherence_historicalCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_adherence_historical"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_adherence_historicalCmd)
}

func Cmdworkforcemanagement_adherence_historical() *cobra.Command {
	return workforcemanagement_adherence_historicalCmd
}
