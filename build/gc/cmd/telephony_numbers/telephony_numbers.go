package telephony_numbers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("telephony_numbers", "SWAGGER_OVERRIDE_/api/v2/telephony/numbers")
	telephony_numbersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_numbers"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_numbersCmd)
}

func Cmdtelephony_numbers() *cobra.Command {
	return telephony_numbersCmd
}
