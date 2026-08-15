package dependencies

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("dependencies", "SWAGGER_OVERRIDE_/api/v2/dependencies")
	dependenciesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("dependencies"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(dependenciesCmd)
}

func Cmddependencies() *cobra.Command {
	return dependenciesCmd
}
