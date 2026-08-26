package externalcontacts_graphs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("externalcontacts_graphs", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/graphs")
	externalcontacts_graphsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_graphs"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_graphsCmd)
}

func Cmdexternalcontacts_graphs() *cobra.Command {
	return externalcontacts_graphsCmd
}
