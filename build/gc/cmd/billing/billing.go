package billing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("billing", "SWAGGER_OVERRIDE_/api/v2/billing")
	billingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("billing"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(billingCmd)
}

func Cmdbilling() *cobra.Command {
	return billingCmd
}
