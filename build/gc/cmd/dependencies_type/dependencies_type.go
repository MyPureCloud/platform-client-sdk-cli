package dependencies_type

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("dependencies_type", "SWAGGER_OVERRIDE_/api/v2/dependencies/type")
	dependencies_typeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("dependencies_type"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(dependencies_typeCmd)
}

func Cmddependencies_type() *cobra.Command {
	return dependencies_typeCmd
}
