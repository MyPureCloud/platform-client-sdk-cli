package outbound_contactlists_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("outbound_contactlists_bulk", "SWAGGER_OVERRIDE_/api/v2/outbound/contactlists/bulk")
	outbound_contactlists_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_contactlists_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_contactlists_bulkCmd)
}

func Cmdoutbound_contactlists_bulk() *cobra.Command {
	return outbound_contactlists_bulkCmd
}
