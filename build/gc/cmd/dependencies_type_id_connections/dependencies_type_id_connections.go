package dependencies_type_id_connections

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("dependencies_type_id_connections", "SWAGGER_OVERRIDE_/api/v2/dependencies/type/{entityType}/id/{entityId}/connections")
	dependencies_type_id_connectionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("dependencies_type_id_connections"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(dependencies_type_id_connectionsCmd)
}

func Cmddependencies_type_id_connections() *cobra.Command {
	return dependencies_type_id_connectionsCmd
}
