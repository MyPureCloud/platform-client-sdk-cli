package orgauthorization

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("orgauthorization", "SWAGGER_OVERRIDE_/api/v2/orgauthorization")
	orgauthorizationCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("orgauthorization"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(orgauthorizationCmd)
}

func Cmdorgauthorization() *cobra.Command {
	return orgauthorizationCmd
}
