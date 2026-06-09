package telephony_organization

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("telephony_organization", "SWAGGER_OVERRIDE_/api/v2/telephony/organization")
	telephony_organizationCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_organization"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_organizationCmd)
}

func Cmdtelephony_organization() *cobra.Command {
	return telephony_organizationCmd
}
