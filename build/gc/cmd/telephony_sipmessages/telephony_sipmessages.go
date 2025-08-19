package telephony_sipmessages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("telephony_sipmessages", "SWAGGER_OVERRIDE_/api/v2/telephony/sipmessages")
	telephony_sipmessagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_sipmessages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_sipmessagesCmd)
}

func Cmdtelephony_sipmessages() *cobra.Command {
	return telephony_sipmessagesCmd
}
