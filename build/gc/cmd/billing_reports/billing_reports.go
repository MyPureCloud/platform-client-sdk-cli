package billing_reports

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("billing_reports", "SWAGGER_OVERRIDE_/api/v2/billing/reports")
	billing_reportsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("billing_reports"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(billing_reportsCmd)
}

func Cmdbilling_reports() *cobra.Command {
	return billing_reportsCmd
}
