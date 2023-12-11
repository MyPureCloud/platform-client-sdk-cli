package outbound_contactlisttemplates_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("outbound_contactlisttemplates_bulk", "SWAGGER_OVERRIDE_/api/v2/outbound/contactlisttemplates/bulk")
	outbound_contactlisttemplates_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_contactlisttemplates_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_contactlisttemplates_bulkCmd)
}

func Cmdoutbound_contactlisttemplates_bulk() *cobra.Command {
	return outbound_contactlisttemplates_bulkCmd
}
