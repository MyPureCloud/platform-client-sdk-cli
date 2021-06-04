package documentationfile_gkn

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("documentationfile_gkn", "SWAGGER_OVERRIDE_/api/v2/documentation/gkn")
	documentationfile_gknCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("documentationfile_gkn"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(documentationfile_gknCmd)
}

func Cmddocumentationfile_gkn() *cobra.Command {
	return documentationfile_gknCmd
}
