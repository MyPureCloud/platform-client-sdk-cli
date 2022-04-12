package externalcontacts_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("externalcontacts_bulk", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/bulk/contacts")
	externalcontacts_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_bulkCmd)
}

func Cmdexternalcontacts_bulk() *cobra.Command {
	return externalcontacts_bulkCmd
}
