package organizations_limits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("organizations_limits", "SWAGGER_OVERRIDE_/api/v2/organizations/limits")
	organizations_limitsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("organizations_limits"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(organizations_limitsCmd)
}

func Cmdorganizations_limits() *cobra.Command {
	return organizations_limitsCmd
}
