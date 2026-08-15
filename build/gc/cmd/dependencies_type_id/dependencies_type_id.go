package dependencies_type_id

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("dependencies_type_id", "SWAGGER_OVERRIDE_/api/v2/dependencies/type/{entityType}/id")
	dependencies_type_idCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("dependencies_type_id"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(dependencies_type_idCmd)
}

func Cmddependencies_type_id() *cobra.Command {
	return dependencies_type_idCmd
}
