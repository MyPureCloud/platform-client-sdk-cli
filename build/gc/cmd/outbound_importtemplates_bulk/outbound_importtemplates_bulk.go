package outbound_importtemplates_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("outbound_importtemplates_bulk", "SWAGGER_OVERRIDE_/api/v2/outbound/importtemplates/bulk")
	outbound_importtemplates_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_importtemplates_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_importtemplates_bulkCmd)
}

func Cmdoutbound_importtemplates_bulk() *cobra.Command {
	return outbound_importtemplates_bulkCmd
}
