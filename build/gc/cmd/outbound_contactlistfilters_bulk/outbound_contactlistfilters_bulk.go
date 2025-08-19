package outbound_contactlistfilters_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("outbound_contactlistfilters_bulk", "SWAGGER_OVERRIDE_/api/v2/outbound/contactlistfilters/bulk")
	outbound_contactlistfilters_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_contactlistfilters_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_contactlistfilters_bulkCmd)
}

func Cmdoutbound_contactlistfilters_bulk() *cobra.Command {
	return outbound_contactlistfilters_bulkCmd
}
