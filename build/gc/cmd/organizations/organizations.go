package organizations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("organizations", "SWAGGER_OVERRIDE_/api/v2/organizations")
	organizationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("organizations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(organizationsCmd)
}

func Cmdorganizations() *cobra.Command {
	return organizationsCmd
}
